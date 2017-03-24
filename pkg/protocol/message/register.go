package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Register{}

// Register will register the supplied Devices with the inventory.
// It is used by drivers to let the system know about devices.
// It is broadcasted at driver startup and unicasted to any
// device requesting the information using a Discovery message.

func (m *Register) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *Register) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *Register) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *Register) Finalize() {
	t := Type_register
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (Register) message() {}

// Used to make sure the interface is met
var _ Message = &Unregister{}

// Unregister will unregister the supplied Devices from the inventory.
// It is used by drivers to let the system know devices have disappeared.
// It is also send at shutdown of a driver.

func (m *Unregister) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *Unregister) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *Unregister) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *Unregister) Finalize() {
	t := Type_unregister
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (Unregister) message() {}
