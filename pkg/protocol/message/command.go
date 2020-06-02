package message

// Used to make sure the interface is met
var _ Message = &Command{}

// Type returns the type of this message
func (m *Command) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *Command) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *Command) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (Command) message() {}

// Used to make sure the interface is met
var _ Message = &CommandReply{}

// Type returns the type of this message
func (m *CommandReply) Type() Type {
	return Type(m.Header.GetType())
}

// From returns the sender of this message
func (m *CommandReply) From() string {
	return m.Header.GetFrom()
}

// To returns the destination of this message
func (m *CommandReply) To() string {
	return m.Header.GetTo()
}

// message is an empty method to comply to the interface Message
func (CommandReply) message() {}
