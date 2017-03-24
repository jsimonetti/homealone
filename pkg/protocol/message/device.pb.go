// Code generated by protoc-gen-go.
// source: device.proto
// DO NOT EDIT!

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Device holds information about a single device.
// A driver can supply multiple devices.
// UUID should be created deterministically in a way
// so that each startup the device gets the same uuid. (uuid.NewV5() helps here)
type Device struct {
	// ID is the identifier of this device
	ID               *string      `protobuf:"bytes,1,req,name=ID,json=iD" json:"ID,omitempty"`
	Owner            *string      `protobuf:"bytes,2,req,name=Owner,json=owner" json:"Owner,omitempty"`
	Name             *string      `protobuf:"bytes,3,req,name=Name,json=name" json:"Name,omitempty"`
	Components       []*Component `protobuf:"bytes,4,rep,name=Components,json=components" json:"Components,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Device) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *Device) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Device) GetComponents() []*Component {
	if m != nil {
		return m.Components
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "message.Device")
}

func init() { proto.RegisterFile("device.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x95, 0x12, 0x48, 0xce, 0xcf, 0x2d, 0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0x86, 0x48, 0x29,
	0x95, 0x71, 0xb1, 0xb9, 0x80, 0x95, 0x0a, 0xf1, 0x71, 0x31, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30,
	0x69, 0x70, 0x06, 0x31, 0x65, 0xba, 0x08, 0x89, 0x70, 0xb1, 0xfa, 0x97, 0xe7, 0xa5, 0x16, 0x49,
	0x30, 0x81, 0x85, 0x58, 0xf3, 0x41, 0x1c, 0x21, 0x21, 0x2e, 0x16, 0xbf, 0xc4, 0xdc, 0x54, 0x09,
	0x66, 0xb0, 0x20, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x90, 0x11, 0x17, 0x97, 0x33, 0xdc, 0x5c, 0x09,
	0x16, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x21, 0x3d, 0xa8, 0x9d, 0x7a, 0x70, 0xa9, 0x20, 0x2e, 0x84,
	0xed, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0xd0, 0xf4, 0xd6, 0xa1, 0x00, 0x00, 0x00,
}
