package message

// Used to make sure the interface is met
var _ Message = &Discover{}

func (m *Discover) Type() Type {
	return Type(m.Header.GetType())
}

func (m *Discover) From() string {
	return m.Header.GetFrom()
}

func (m *Discover) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Discover) message() {}
