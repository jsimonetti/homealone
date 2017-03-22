package message

import (
	"github.com/jsimonetti/homealone/protocol/device"
)

// Used to make sure the interface is met
var _ Message = &Register{}

// Register will register the supplied Devices with the inventory.
// It is used by drivers to let the system know about devices.
// It is broadcasted at driver startup and unicasted to any
// device requesting the information using a Discovery message.
type Register struct {
	Header
	Name    string
	Devices []device.Device
}

// Finalize will finish the object before marshalling
func (m *Register) Finalize() {
	m.Header.MType = TypeRegister
}

// message is an empty method to comply to the interface Message
func (Register) message() {}

// Used to make sure the interface is met
var _ Message = &Unregister{}

// Unregister will unregister the supplied Devices from the inventory.
// It is used by drivers to let the system know devices have disappeared.
// It is also send at shutdown of a driver.
type Unregister struct {
	Header
	Name    string
	Devices []device.Device
}

// Finalize will finish the object before marshalling
func (m *Unregister) Finalize() {
	m.Header.MType = TypeUnregister
}

// message is an empty method to comply to the interface Message
func (Unregister) message() {}
