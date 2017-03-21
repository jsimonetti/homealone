package message

import (
	"github.com/jsimonetti/homealone/protocol/device"
)

// Used to make sure the interface is met
var _ Message = &Discover{}

// Discover is a message used by the inventory to discover
// all available devices in the network.
type Discover struct {
	Header
}

// Finalize will finish the object before marshalling
func (m *Discover) Finalize() {
	m.Header.MType = TypeDiscover
}

// message is an empty method to comply to the interface Message
func (Discover) message() {}

// Used to make sure the interface is met
var _ Message = &DiscoverReply{}

// DiscoverReply is similar to a Register message except
// it is send in reply to a Discover message.
// This allows for a stateless system.
type DiscoverReply struct {
	Header
	Name    string
	Devices []device.Device
}

// Finalize will finish the object before marshalling
func (m *DiscoverReply) Finalize() {
	m.Header.MType = TypeDiscoverReply
}

// message is an empty method to comply to the interface Message
func (DiscoverReply) message() {}
