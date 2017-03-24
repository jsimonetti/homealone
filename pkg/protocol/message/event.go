package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Event{}

func (m *Event) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *Event) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *Event) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *Event) Finalize() {
	t := Type_event
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (Event) message() {}
