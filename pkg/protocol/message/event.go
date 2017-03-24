package message

// Used to make sure the interface is met
var _ Message = &Event{}

func (m *Event) Type() Type {
	return Type(m.Header.GetType())
}

func (m *Event) From() string {
	return m.Header.GetFrom()
}

func (m *Event) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Event) message() {}
