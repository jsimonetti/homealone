package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Command{}

// Command is used to send a command to a device on the network.
// The destination is the device UUID
type Command struct {
	Header
	ID          uuid.UUID
	Destination uuid.UUID
	Component   string
	Op          string
}

// Finalize will finish the object before marshalling
func (m *Command) Finalize() {
	m.Header.MType = TypeCommand
}

// message is an empty method to comply to the interface Message
func (Command) message() {}

// Used to make sure the interface is met
var _ Message = &CommandReply{}

// CommandReply is sent in response to a Command.
// It contains the result of a command and an optional message
type CommandReply struct {
	Header
	InReplyTo uuid.UUID
	Result    CommandResult
	Message   string
}

// Finalize will finish the object before marshalling
func (m *CommandReply) Finalize() {
	m.Header.MType = TypeCommandReply
}

// message is an empty method to comply to the interface Message
func (CommandReply) message() {}

//go:generate stringer -type=CommandResult

// CommandResult is used to signal the sending partner about the result
// of this command
type CommandResult uint8

const (
	// CommandSyncAck is used to acknowledge receipt of the command and
	// successful execution
	CommandSyncAck CommandResult = iota
	// CommandSyncFail is used to acknowledge receipt of the command and
	// unsuccessful execution
	CommandSyncFail
	// CommandAsyncAck is used to acknowledge receipt of the command but
	// immediate result is not available. The caller should monitor the
	// command queue to wait for the async result
	CommandAsyncAck
	// CommandAsyncFail is used to signal failure on an async command.
	// Message should be consulted to see the reason.
	CommandAsyncFail
	// CommandAsyncSuccess is used to signal success on an async command
	CommandAsyncSuccess
	// CommandError is used to signal a failure with executing the command.
	// Message should be consulted to see the reason.
	CommandError
)
