package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Discover{}

func (m *Discover) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *Discover) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *Discover) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *Discover) Finalize() {
	m.Header.Mtype = Type_discover.Enum()
}

// message is an empty method to comply to the interface Message
func (Discover) message() {}
