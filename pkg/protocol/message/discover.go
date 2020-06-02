package message

// Used to make sure the interface is met
var _ Message = &Discover{}

// Type returns the type of this message
func (m *Discover) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *Discover) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *Discover) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Discover) message() {}
