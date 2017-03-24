package message

import (
	"github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
)

type Message interface {
	proto.Message
	Type() Type
	From() uuid.UUID
	To() uuid.UUID
	Finalize()
	message()
}

// Handler is used by the app framework to register a callback function
type Handler func(m Message) error
