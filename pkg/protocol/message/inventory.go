package message

// Used to make sure the interface is met
var _ Message = &Inventory{}

func (m *Inventory) Type() Type {
	return Type(m.Header.GetType())
}

func (m *Inventory) From() string {
	return m.Header.GetFrom()
}

func (m *Inventory) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Inventory) message() {}

// Used to make sure the interface is met
var _ Message = &InventoryReply{}

func (m *InventoryReply) Type() Type {
	return Type(m.Header.GetType())
}

func (m *InventoryReply) From() string {
	return m.Header.GetFrom()
}

func (m *InventoryReply) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (InventoryReply) message() {}
