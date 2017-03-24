// Code generated by protoc-gen-go.
// source: type.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	type.proto
	header.proto
	discover.proto
	command.proto
	event.proto
	register.proto
	device.proto
	components.proto
	inventory.proto

It has these top-level messages:
	Header
	Discover
	Command
	CommandReply
	Event
	Register
	Unregister
	Device
	Components
	Stat
	Toggle
	Slider
	Rotary
	Inventory
	InventoryReply
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Type is the type of the message. This is used for (un)marshalling.
type Type int32

const (
	// TypeDiscover is sent by the inventory app to discover all devices.
	// Drivers should respond to this using a Register message.
	// It is always broadcasted on the inventory queue.
	Type_discover Type = 1
	// TypeRegister is sent by drivers to register devices with the inventory.
	// Drivers using the App framework automatically do this at startup, and
	// in respone to a Discover message. It is broadcasted on startup and
	// unicasted in reply to a discover on the inventory queue.
	Type_register Type = 2
	// TypeUnregister will unregister a device from the inventory.
	// A driver may unregister a device if it is removed or otherwise unavailable
	// for commands. A temporary failure should not result in an unregistered
	// device. Drivers using the App framework will automatically unregister
	// all their (current) devices at shutdown. It is always broadcasted on the
	// inventory queue.
	Type_unregister Type = 3
	// TypeInventory is used to retrieve an inventory of devices from the
	// inventory app. It results in a InventoryReply from the inventory. Only
	// the inventory should respond to this type of message. It is always
	// broadcasted on the inventory queue.
	Type_inventory Type = 4
	// TypeInventoryReply is used by the inventory app to send an inventory
	// of all devices in the network to the requester. It is unicasted to the
	// requester on the inventory queue.
	Type_inventoryReply Type = 5
	// TypeCommand is used to send a command to a device. It can be broadcasted
	// or unicasted on the command queue. In case it is unicasted to a specific
	// driver, that driver should always respond. In case of a broadcast, only
	// the driver holding the target device should respond, but a respons is not
	// guaranteed or required.
	Type_command Type = 6
	// TypeCommandReply is a respons to a TypeCommand. It relays the success of
	// the command. It is always unicasted directly to the requester on the
	// command queue.
	Type_commandReply Type = 7
	// TypeEvent is used by drivers or apps to send events about state changes,
	// sensor changes and/or other events. It is always broadcasted on the
	// event queue.
	Type_event Type = 8
)

var Type_name = map[int32]string{
	1: "discover",
	2: "register",
	3: "unregister",
	4: "inventory",
	5: "inventoryReply",
	6: "command",
	7: "commandReply",
	8: "event",
}
var Type_value = map[string]int32{
	"discover":       1,
	"register":       2,
	"unregister":     3,
	"inventory":      4,
	"inventoryReply": 5,
	"command":        6,
	"commandReply":   7,
	"event":          8,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("message.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("type.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcb, 0x5f, 0x0a, 0x82, 0x40,
	0x14, 0x85, 0x71, 0x2a, 0x4d, 0x3d, 0x99, 0x5c, 0xee, 0x32, 0x7a, 0x68, 0x31, 0xd1, 0x06, 0x44,
	0x0f, 0x32, 0xd0, 0xfc, 0x61, 0x66, 0x12, 0xe6, 0xa9, 0xad, 0x87, 0x14, 0x3e, 0xfe, 0xf8, 0xf8,
	0x80, 0x5c, 0x02, 0xef, 0x21, 0xfa, 0xec, 0xb5, 0xb1, 0x4c, 0x69, 0x5c, 0x78, 0xfb, 0xa0, 0x7a,
	0x96, 0x40, 0xed, 0xd1, 0xce, 0x26, 0x4d, 0x7e, 0x65, 0x94, 0xc3, 0xa6, 0xc8, 0xc5, 0xa4, 0xcc,
	0x28, 0x47, 0x1d, 0x80, 0xb7, 0xdb, 0x7d, 0xd2, 0x2b, 0x3a, 0xe3, 0x56, 0xba, 0xec, 0x63, 0x91,
	0x4a, 0x15, 0xc3, 0xce, 0x07, 0xc3, 0xab, 0x48, 0xad, 0x17, 0x34, 0x93, 0xb7, 0x76, 0x74, 0xb3,
	0x9c, 0x55, 0xd0, 0xff, 0xf1, 0xcb, 0x8d, 0x76, 0xa8, 0xb9, 0x1d, 0xd2, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0x82, 0x01, 0xcc, 0x96, 0x00, 0x00, 0x00,
}
