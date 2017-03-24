package protocol

import (
	"bytes"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
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
	m = finalizeMessage(m)

	data, err := proto.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "marshal failed")
	}

	b := append(version[:], uint8(m.Type()))
	return append(b, data...), nil
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

	m, err = decodeMessage(t, buf.Bytes())
	return m, errors.Wrap(err, "decodeMessage failed")
}

// decodeMessage will decode the bytes from the buffer according to the type.
// The remaining bytes on the buffer are decoded into the appropriate message.
func decodeMessage(t message.Type, b []byte) (m message.Message, err error) {

	switch t {
	case message.Type_discover:
		msg := &message.Discover{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_register:
		msg := &message.Register{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_unregister:
		msg := &message.Unregister{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_inventory:
		msg := &message.Inventory{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_inventoryReply:
		msg := &message.InventoryReply{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_command:
		msg := &message.Command{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_commandReply:
		msg := &message.CommandReply{}
		err = proto.Unmarshal(b, msg)
		m = msg

	case message.Type_event:
		msg := &message.Event{}
		err = proto.Unmarshal(b, msg)
		m = msg

	default:
		return nil, fmt.Errorf("unknown messagetype found; type: %s", t.String())
	}

	return m, errors.Wrap(err, "proto.Decode failed")
}

// Payload is the interface for Publish payloads.
type Payload interface {
	// WritePayload writes the payload data to w.
	WritePayload(w io.Writer) error
}

// decodeMessage will decode the bytes from the buffer according to the type.
// The remaining bytes on the buffer are decoded into the appropriate message.
func finalizeMessage(msg message.Message) (m message.Message) {

	switch msg := msg.(type) {
	case *message.Discover:
		msg.Header.Type = message.Type_discover.Enum()

	case *message.Register:
		msg.Header.Type = message.Type_register.Enum()

	case *message.Unregister:
		msg.Header.Type = message.Type_unregister.Enum()

	case *message.Inventory:
		msg.Header.Type = message.Type_inventory.Enum()

	case *message.InventoryReply:
		msg.Header.Type = message.Type_inventoryReply.Enum()

	case *message.Command:
		msg.Header.Type = message.Type_command.Enum()

	case *message.CommandReply:
		msg.Header.Type = message.Type_commandReply.Enum()

	case *message.Event:
		msg.Header.Type = message.Type_event.Enum()

	default:
		return nil
	}
	m = msg

	return m
}
