package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Inventory{}

// Inventory will ask for an inventory of devices.
// It is broadcasted across the network

func (m *Inventory) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *Inventory) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *Inventory) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *Inventory) Finalize() {
	t := Type_inventory
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (Inventory) message() {}

// Used to make sure the interface is met
var _ Message = &InventoryReply{}

// InventoryReply is a reply to an inventory request.
// It is unicasted to the requester and holds all
// devices currently in the inventory

func (m *InventoryReply) Type() Type {
	return Type(m.Header.GetMtype())
}

func (m *InventoryReply) From() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMfrom())
	return id
}

func (m *InventoryReply) To() uuid.UUID {
	id, _ := uuid.FromBytes(m.Header.GetMto())
	return id
}

// Finalize will finish the object before marshalling
func (m *InventoryReply) Finalize() {
	t := Type_inventoryReply
	m.Header.Mtype = &t
}

// message is an empty method to comply to the interface Message
func (InventoryReply) message() {}
