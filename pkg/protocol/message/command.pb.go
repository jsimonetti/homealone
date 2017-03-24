// Code generated by protoc-gen-go.
// source: command.proto
// DO NOT EDIT!

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// CommandResult is used to signal the sending partner about the result
// of this command
type CommandResult int32

const (
	// CommandSyncAck is used to acknowledge receipt of the command and
	// successful execution
	CommandResult_SyncAck CommandResult = 1
	// CommandSyncFail is used to acknowledge receipt of the command and
	// unsuccessful execution
	CommandResult_SyncFail CommandResult = 2
	// CommandAsyncAck is used to acknowledge receipt of the command but
	// immediate result is not available. The caller should monitor the
	// command queue to wait for the async result
	CommandResult_AsyncAck CommandResult = 3
	// CommandAsyncFail is used to signal failure on an async command.
	// Message should be consulted to see the reason.
	CommandResult_AsyncFail CommandResult = 4
	// CommandAsyncSuccess is used to signal success on an async command
	CommandResult_AsyncSuccess CommandResult = 5
	// CommandError is used to signal a failure with executing the command.
	// Message should be consulted to see the reason.
	CommandResult_Error CommandResult = 6
)

var CommandResult_name = map[int32]string{
	1: "SyncAck",
	2: "SyncFail",
	3: "AsyncAck",
	4: "AsyncFail",
	5: "AsyncSuccess",
	6: "Error",
}
var CommandResult_value = map[string]int32{
	"SyncAck":      1,
	"SyncFail":     2,
	"AsyncAck":     3,
	"AsyncFail":    4,
	"AsyncSuccess": 5,
	"Error":        6,
}

func (x CommandResult) Enum() *CommandResult {
	p := new(CommandResult)
	*p = x
	return p
}
func (x CommandResult) String() string {
	return proto.EnumName(CommandResult_name, int32(x))
}
func (x *CommandResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandResult_value, data, "CommandResult")
	if err != nil {
		return err
	}
	*x = CommandResult(value)
	return nil
}
func (CommandResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// Command is used to send a command to a device on the network.
// The destination is the device UUID
type Command struct {
	Header           *Header `protobuf:"bytes,1,req,name=Header,json=header" json:"Header,omitempty"`
	ID               *string `protobuf:"bytes,2,req,name=ID,json=iD" json:"ID,omitempty"`
	Destination      *string `protobuf:"bytes,3,req,name=Destination,json=destination" json:"Destination,omitempty"`
	Component        *string `protobuf:"bytes,4,req,name=Component,json=component" json:"Component,omitempty"`
	Op               *string `protobuf:"bytes,5,req,name=Op,json=op" json:"Op,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Command) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Command) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *Command) GetDestination() string {
	if m != nil && m.Destination != nil {
		return *m.Destination
	}
	return ""
}

func (m *Command) GetComponent() string {
	if m != nil && m.Component != nil {
		return *m.Component
	}
	return ""
}

func (m *Command) GetOp() string {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return ""
}

// CommandReply is sent in response to a Command.
// It contains the result of a command and an optional message
type CommandReply struct {
	Header           *Header        `protobuf:"bytes,1,req,name=Header,json=header" json:"Header,omitempty"`
	InReplyTo        *string        `protobuf:"bytes,2,req,name=InReplyTo,json=inReplyTo" json:"InReplyTo,omitempty"`
	Result           *CommandResult `protobuf:"varint,3,req,name=Result,json=result,enum=message.CommandResult" json:"Result,omitempty"`
	Message          *string        `protobuf:"bytes,4,opt,name=Message,json=message" json:"Message,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *CommandReply) Reset()                    { *m = CommandReply{} }
func (m *CommandReply) String() string            { return proto.CompactTextString(m) }
func (*CommandReply) ProtoMessage()               {}
func (*CommandReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *CommandReply) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommandReply) GetInReplyTo() string {
	if m != nil && m.InReplyTo != nil {
		return *m.InReplyTo
	}
	return ""
}

func (m *CommandReply) GetResult() CommandResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return CommandResult_SyncAck
}

func (m *CommandReply) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "message.Command")
	proto.RegisterType((*CommandReply)(nil), "message.CommandReply")
	proto.RegisterEnum("message.CommandResult", CommandResult_name, CommandResult_value)
}

func init() { proto.RegisterFile("command.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x6e, 0xb3, 0x30,
	0x14, 0x84, 0x85, 0x93, 0xc0, 0xef, 0x07, 0xc9, 0x8f, 0xde, 0xa2, 0xb2, 0xaa, 0x2c, 0x50, 0x36,
	0x8d, 0xba, 0x60, 0x91, 0x1b, 0x44, 0xa1, 0x55, 0xb3, 0xa8, 0x2a, 0x39, 0xbd, 0x00, 0x32, 0x56,
	0x8b, 0x0a, 0x36, 0xb2, 0xc9, 0x22, 0xd7, 0xe8, 0x05, 0x7a, 0xd5, 0x0a, 0x1b, 0x52, 0x75, 0xd7,
	0x15, 0x9a, 0xf9, 0xde, 0x0c, 0x23, 0xc3, 0x52, 0xe8, 0xb6, 0x2d, 0x55, 0x95, 0x77, 0x46, 0xf7,
	0x1a, 0xa3, 0x56, 0x5a, 0x5b, 0xbe, 0xc9, 0xdb, 0xe4, 0x5d, 0x96, 0x95, 0x34, 0xde, 0xde, 0x7c,
	0x06, 0x10, 0x1d, 0xfc, 0x21, 0xde, 0x41, 0xf8, 0xe4, 0x18, 0x0b, 0x32, 0xb2, 0x8d, 0x77, 0xff,
	0xf3, 0x31, 0x93, 0x7b, 0x9b, 0x87, 0x3e, 0x8a, 0x2b, 0x20, 0xc7, 0x82, 0x91, 0x8c, 0x6c, 0x29,
	0x27, 0x75, 0x81, 0x19, 0xc4, 0x85, 0xb4, 0x7d, 0xad, 0xca, 0xbe, 0xd6, 0x8a, 0xcd, 0x1c, 0x88,
	0xab, 0x1f, 0x0b, 0xd7, 0x40, 0x0f, 0xba, 0xed, 0xb4, 0x92, 0xaa, 0x67, 0x73, 0xc7, 0xa9, 0x98,
	0x8c, 0xa1, 0xef, 0xa5, 0x63, 0x0b, 0xdf, 0xa7, 0xbb, 0xcd, 0x57, 0x00, 0xc9, 0x38, 0x8a, 0xcb,
	0xae, 0xb9, 0xfc, 0x7d, 0xd9, 0x1a, 0xe8, 0x51, 0xb9, 0xcc, 0xab, 0x1e, 0x07, 0xd2, 0x7a, 0x32,
	0x30, 0x87, 0x90, 0x4b, 0x7b, 0x6e, 0x7a, 0x37, 0x71, 0xb5, 0xbb, 0xb9, 0xd6, 0x5c, 0xff, 0x36,
	0x50, 0x1e, 0x1a, 0xf7, 0x45, 0x06, 0xd1, 0xb3, 0x3f, 0x60, 0xf3, 0x2c, 0xd8, 0x52, 0x3e, 0x3d,
	0xe2, 0x7d, 0x05, 0xcb, 0x5f, 0x11, 0x8c, 0x21, 0x3a, 0x5d, 0x94, 0xd8, 0x8b, 0x8f, 0x34, 0xc0,
	0x04, 0xfe, 0x0d, 0xe2, 0xb1, 0xac, 0x9b, 0x94, 0x0c, 0x6a, 0x6f, 0x47, 0x36, 0xc3, 0x25, 0x50,
	0xa7, 0x1c, 0x9c, 0x63, 0x0a, 0x89, 0x93, 0xa7, 0xb3, 0x10, 0xd2, 0xda, 0x74, 0x81, 0x14, 0x16,
	0x0f, 0xc6, 0x68, 0x93, 0x86, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xd0, 0x0f, 0x7a, 0xc3,
	0x01, 0x00, 0x00,
}
