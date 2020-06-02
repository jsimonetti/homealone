package message

// Used to make sure the interface is met
var _ Message = &Inventory{}

// Type returns the type of this message
func (m *Inventory) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *Inventory) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *Inventory) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Inventory) message() {}

// Used to make sure the interface is met
var _ Message = &InventoryReply{}

// Type returns the type of this message
func (m *InventoryReply) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *InventoryReply) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *InventoryReply) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (InventoryReply) message() {}
