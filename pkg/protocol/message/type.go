package message

//go:generate stringer -type=Type

// Type is the type of the message. This is used for (un)marshalling.
type Type uint8

// These constants define the different message types.
const (
	TypeDiscover Type = iota
	TypeRegister
	TypeUnregister
	TypeInventory
	TypeInventoryReply
	TypeCommand
	TypeCommandReply
)
