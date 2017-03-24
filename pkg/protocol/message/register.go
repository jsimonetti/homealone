package message

// Used to make sure the interface is met
var _ Message = &Register{}

// Type returns the type of this message
func (m *Register) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *Register) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *Register) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Register) message() {}

// Used to make sure the interface is met
var _ Message = &Unregister{}

// Type returns the type of this message
func (m *Unregister) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *Unregister) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *Unregister) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Unregister) message() {}
