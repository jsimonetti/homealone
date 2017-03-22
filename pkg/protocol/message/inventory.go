package message

import (
	"github.com/jsimonetti/homealone/pkg/protocol/device"
)

// Used to make sure the interface is met
var _ Message = &Inventory{}

// Inventory will ask for an inventory of devices.
// It is broadcasted across the network
type Inventory struct {
	Header
}

// Finalize will finish the object before marshalling
func (m *Inventory) Finalize() {
	m.Header.MType = TypeInventory
}

// message is an empty method to comply to the interface Message
func (Inventory) message() {}

// Used to make sure the interface is met
var _ Message = &InventoryReply{}

// InventoryReply is a reply to an inventory request.
// It is unicasted to the requester and holds all
// devices currently in the inventory
type InventoryReply struct {
	Header
	Devices []device.Device
}

// Finalize will finish the object before marshalling
func (m *InventoryReply) Finalize() {
	m.Header.MType = TypeInventoryReply
}

// message is an empty method to comply to the interface Message
func (InventoryReply) message() {}
