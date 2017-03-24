package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Command{}

func (m *Command) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *Command) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *Command) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *Command) Finalize() {
	t := Type_command
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (Command) message() {}

// Used to make sure the interface is met
var _ Message = &CommandReply{}

func (m *CommandReply) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *CommandReply) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *CommandReply) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *CommandReply) Finalize() {
	t := Type_commandReply
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (CommandReply) message() {}
