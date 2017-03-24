package message

//go:generate stringer -type=Type

// Type is the type of the message. This is used for (un)marshalling.
type Type uint8

// These constants define the different message types.
const (
	// TypeDiscover is sent by the inventory app to discover all devices.
	// Drivers should respond to this using a Register message.
	// It is always broadcasted on the inventory queue.
	TypeDiscover Type = iota
	// TypeRegister is sent by drivers to register devices with the inventory.
	// Drivers using the App framework automatically do this at startup, and
	// in respone to a Discover message. It is broadcasted on startup and unicasted in
	// reply to a discover on the inventory queue.
	TypeRegister
	// TypeUnregister will unregister a device from the inventory.
	// A driver may unregister a device if it is removed or otherwise unavailable for commands.
	// A temporary failure should not result in an unregistered device. Drivers using the App
	// framework will automatically unregister all their (current) devices at shutdown.
	// It is always broadcasted on the inventory queue.
	TypeUnregister
	// TypeInventory is used to retrieve an inventory of devices from the inventory app.
	// It results in a InventoryReply from the inventory. Only the inventory should respond to this type
	// of message. It is always broadcasted on the inventory queue.
	TypeInventory
	// TypeInventoryReply is used by the inventory app to send an inventory of all devices in the network
	// to the requester. It is unicasted to the requester on the inventory queue.
	TypeInventoryReply
	// TypeCommand is used to send a command to a device. It can be broadcasted or unicasted on the command
	// queue. In case it is unicasted to a specific driver, that driver should always respond.
	// In case of a broadcast, only the driver holding the target device should respond, but a respons is
	// not guaranteed or required.
	TypeCommand
	// TypeCommandReply is a respons to a TypeCommand. It relays the success of the command. It is always
	// unicasted directly to the requester on the command queue.
	TypeCommandReply
	// TypeEvent is used by drivers or apps to send events about state changes, sensor changes and/or other events.
	// It is always broadcasted on the even queue.
	TypeEvent
)
