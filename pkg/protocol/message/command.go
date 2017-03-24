package message

// Used to make sure the interface is met
var _ Message = &Command{}

func (m *Command) Type() Type {
	return Type(m.Header.GetType())
}

func (m *Command) From() string {
	return m.Header.GetFrom()
}

func (m *Command) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Command) message() {}

// Used to make sure the interface is met
var _ Message = &CommandReply{}

func (m *CommandReply) Type() Type {
	return Type(m.Header.GetType())
}

func (m *CommandReply) From() string {
	return m.Header.GetFrom()
}

func (m *CommandReply) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (CommandReply) message() {}
