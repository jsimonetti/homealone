// Package protocol defines the protocol used to talk to the hub
package protocol

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/jsimonetti/homealone/pkg/protocol/message"
)

// version is the current version of the protocol
var version = [8]byte{'h', 'a', '-', '1'}

// errVersionMismatch is thrown when the protocol versions do not match
const errVersionMismatch = protoError("Protocol version mismatch")

type protoError string

func (e protoError) Error() string { return string(e) }

// Marshal will marshal a Message into bytes
// The version and message type are prepended onto the binary data
func Marshal(m message.Message) ([]byte, error) {
	m.Finalize()
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, errors.Wrap(err, "marshal failed")
	}

	b := append(version[:], uint8(m.Type()))
	return append(b, buf.Bytes()...), nil
}

// Unmarshal will take the bytes from an mqtt.Payload and read them into a Buffer
// The version is checked with the current version to see if it matches
func Unmarshal(p Payload) (m message.Message, err error) {
	buf := &bytes.Buffer{}
	p.WritePayload(buf)

	v := make([]byte, 8)
	_, err = buf.Read(v)
	if err != nil {
		return nil, errors.Wrap(err, "read version failed")
	}

	if !bytes.Equal(version[:], v) {
		return nil, errVersionMismatch
	}

	var b byte
	b, err = buf.ReadByte()
	if err != nil {
		return nil, errors.Wrap(err, "read message type failed")
	}

	t := message.Type(b)
	m, err = decodeMessage(t, buf)
	return m, errors.Wrap(err, "decodeMessage failed")
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

	case message.TypeRegister:
		msg := &message.Register{}
		err = dec.Decode(&msg)
		m = msg

	case message.TypeUnregister:
		msg := &message.Unregister{}
		err = dec.Decode(&msg)
		m = msg

	case message.TypeInventory:
		msg := &message.Inventory{}
		err = dec.Decode(&msg)
		m = msg

	case message.TypeInventoryReply:
		msg := &message.InventoryReply{}
		err = dec.Decode(&msg)
		m = msg

	default:
		return nil, fmt.Errorf("unknown messagetype found; type: %s", t.String())
	}

	return m, errors.Wrap(err, "gob.Decode failed")
}

// Payload is the interface for Publish payloads.
type Payload interface {
	// Size returns the number of bytes that WritePayload will write.
	Size() int

	// WritePayload writes the payload data to w. Implementations must write
	// Size() bytes of data, but it is *not* required to do so prior to
	// returning. Size() bytes must have been written to w prior to another
	// message being encoded to the underlying connection.
	WritePayload(w io.Writer) error

	// ReadPayload reads the payload data from r (r will EOF at the end of the
	// payload). It is *not* required for r to have been consumed prior to this
	// returning. r must have been consumed completely prior to another message
	// being decoded from the underlying connection.
	ReadPayload(r io.Reader) error
}
