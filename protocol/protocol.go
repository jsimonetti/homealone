// Package protocol defines the protocol used to talk to the hub
package protocol

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/huin/mqtt"
	"github.com/jsimonetti/homealone/protocol/message"
)

// version is the current version of the protocol
var version = [8]byte{'h', 'a', '-', '1'}

// Marshal will marshal a Message into bytes
// The version and message type are prepended onto the binary data
func Marshal(m message.Message) ([]byte, error) {
	m.Finalize()
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}

	b := append(version[:], uint8(m.Type()))
	return append(b, buf.Bytes()...), nil
}

// Unmarshal will take the bytes from an mqtt.Payload and read them into a Buffer
// The version is checked with the current version to see if it matches
func Unmarshal(p mqtt.Payload) (m message.Message, err error) {
	buf := &bytes.Buffer{}
	p.WritePayload(buf)

	v := make([]byte, 8)
	_, err = buf.Read(v)
	if err != nil {
		return
	}

	if !bytes.Equal(version[:], v) {
		return nil, fmt.Errorf("protocol version mismatch; want: '%v', got '%v'", strings.TrimRight(string(version[:]), "\x00"), strings.TrimRight(string(v), "\x00"))
	}

	var b byte
	b, err = buf.ReadByte()
	if err != nil {
		return
	}

	t := message.Type(b)
	m, err = decodeMessage(t, buf)
	return
}

// decodeMessage will decode the bytes from the buffer according to the type.
// The remaining bytes on the buffer are decoded into the appropriate message.
func decodeMessage(t message.Type, buf *bytes.Buffer) (m message.Message, err error) {
	dec := gob.NewDecoder(buf)

	switch t {
	case message.TypeDiscover:
		msg := &message.Discover{}
		err = dec.Decode(&msg)
		m = msg

	case message.TypeDiscoverReply:
		msg := &message.DiscoverReply{}
		err = dec.Decode(&msg)
		m = msg

	case message.TypeRegister:
		msg := &message.Register{}
		err = dec.Decode(&msg)
		m = msg

	case message.TypeUnregister:
		msg := &message.Unregister{}
		err = dec.Decode(&msg)
		m = msg

	default:
		return nil, fmt.Errorf("unknown messagetype found; type: %s", t.String())
	}

	return
}
