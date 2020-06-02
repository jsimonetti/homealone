package message

import (
	"github.com/golang/protobuf/proto"
)

// Message is the interface used for messages
type Message interface {
	proto.Message
	Type() Type
	From() string
	To() string
	message()
}

// Handler is used by the app framework to register a callback function
type Handler func(m Message) error
