package message

// Used to make sure the interface is met
var _ Message = &Event{}

// Type returns the type of this message
func (m *Event) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *Event) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *Event) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Event) message() {}
