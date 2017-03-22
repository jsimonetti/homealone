package message

// Used to make sure the interface is met
var _ Message = &Discover{}

// Discover is a message used by the inventory to discover
// all available devices in the network.
// Drivers should respond with a Register message directed to the source
type Discover struct {
	Header
}

// Finalize will finish the object before marshalling
func (m *Discover) Finalize() {
	m.Header.MType = TypeDiscover
}

// message is an empty method to comply to the interface Message
func (Discover) message() {}
