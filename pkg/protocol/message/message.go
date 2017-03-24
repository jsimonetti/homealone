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
