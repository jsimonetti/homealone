package message

// Used to make sure the interface is met
var _ Message = &Register{}

func (m *Register) Type() Type {
	return Type(m.Header.GetType())
}

func (m *Register) From() string {
	return m.Header.GetFrom()
}

func (m *Register) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Register) message() {}

// Used to make sure the interface is met
var _ Message = &Unregister{}

func (m *Unregister) Type() Type {
	return Type(m.Header.GetType())
}

func (m *Unregister) From() string {
	return m.Header.GetFrom()
}

func (m *Unregister) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Unregister) message() {}
