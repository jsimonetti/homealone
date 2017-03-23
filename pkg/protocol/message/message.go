package message

import (
	uuid "github.com/satori/go.uuid"
)

// Header is the header of a message. It contains the message type and
// the source of the message.
type Header struct {
	MType  Type
	Source uuid.UUID
	For    uuid.UUID
}

// Type will return this messages' type
func (h *Header) Type() Type {
	return h.MType
}

// From will return this messages' source
func (h *Header) From() uuid.UUID {
	return h.Source
}

// To will return this messages' destination
func (h *Header) To() uuid.UUID {
	return h.For
}

// Message is the interface used for this message.
type Message interface {
	Type() Type
	From() uuid.UUID
	To() uuid.UUID
	Finalize()
	message()
}

// Handler is used by the app framework to register a callback function
type Handler func(m Message) error
