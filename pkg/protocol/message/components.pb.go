// Code generated by protoc-gen-go.
// source: components.proto
// DO NOT EDIT!

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Component is an interface to device components
type Component struct {
	// Types that are valid to be assigned to Union:
	//	*Component_Stat
	//	*Component_Toggle
	//	*Component_Slider
	//	*Component_Rotary
	Union            isComponent_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Component) Reset()                    { *m = Component{} }
func (m *Component) String() string            { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()               {}
func (*Component) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type isComponent_Union interface {
	isComponent_Union()
}

type Component_Stat struct {
	Stat *Stat `protobuf:"bytes,1,opt,name=Stat,json=stat,oneof"`
}
type Component_Toggle struct {
	Toggle *Toggle `protobuf:"bytes,2,opt,name=Toggle,json=toggle,oneof"`
}
type Component_Slider struct {
	Slider *Slider `protobuf:"bytes,3,opt,name=Slider,json=slider,oneof"`
}
type Component_Rotary struct {
	Rotary *Rotary `protobuf:"bytes,4,opt,name=Rotary,json=rotary,oneof"`
}

func (*Component_Stat) isComponent_Union()   {}
func (*Component_Toggle) isComponent_Union() {}
func (*Component_Slider) isComponent_Union() {}
func (*Component_Rotary) isComponent_Union() {}

func (m *Component) GetUnion() isComponent_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Component) GetStat() *Stat {
	if x, ok := m.GetUnion().(*Component_Stat); ok {
		return x.Stat
	}
	return nil
}

func (m *Component) GetToggle() *Toggle {
	if x, ok := m.GetUnion().(*Component_Toggle); ok {
		return x.Toggle
	}
	return nil
}

func (m *Component) GetSlider() *Slider {
	if x, ok := m.GetUnion().(*Component_Slider); ok {
		return x.Slider
	}
	return nil
}

func (m *Component) GetRotary() *Rotary {
	if x, ok := m.GetUnion().(*Component_Rotary); ok {
		return x.Rotary
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Component) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Component_OneofMarshaler, _Component_OneofUnmarshaler, _Component_OneofSizer, []interface{}{
		(*Component_Stat)(nil),
		(*Component_Toggle)(nil),
		(*Component_Slider)(nil),
		(*Component_Rotary)(nil),
	}
}

func _Component_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Component)
	// union
	switch x := m.Union.(type) {
	case *Component_Stat:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Stat); err != nil {
			return err
		}
	case *Component_Toggle:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Toggle); err != nil {
			return err
		}
	case *Component_Slider:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Slider); err != nil {
			return err
		}
	case *Component_Rotary:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rotary); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Component.Union has unexpected type %T", x)
	}
	return nil
}

func _Component_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Component)
	switch tag {
	case 1: // union.Stat
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Stat)
		err := b.DecodeMessage(msg)
		m.Union = &Component_Stat{msg}
		return true, err
	case 2: // union.Toggle
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Toggle)
		err := b.DecodeMessage(msg)
		m.Union = &Component_Toggle{msg}
		return true, err
	case 3: // union.Slider
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Slider)
		err := b.DecodeMessage(msg)
		m.Union = &Component_Slider{msg}
		return true, err
	case 4: // union.Rotary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Rotary)
		err := b.DecodeMessage(msg)
		m.Union = &Component_Rotary{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Component_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Component)
	// union
	switch x := m.Union.(type) {
	case *Component_Stat:
		s := proto.Size(x.Stat)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Toggle:
		s := proto.Size(x.Toggle)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Slider:
		s := proto.Size(x.Slider)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Rotary:
		s := proto.Size(x.Rotary)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Stat is a statistic and can be any value you desire
type Stat struct {
	Name             *string `protobuf:"bytes,1,req,name=Name,json=name" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Stat) Reset()                    { *m = Stat{} }
func (m *Stat) String() string            { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()               {}
func (*Stat) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *Stat) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// Toggle is a simple binary toggle
type Toggle struct {
	Name             *string `protobuf:"bytes,1,req,name=Name,json=name" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Toggle) Reset()                    { *m = Toggle{} }
func (m *Toggle) String() string            { return proto.CompactTextString(m) }
func (*Toggle) ProtoMessage()               {}
func (*Toggle) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Toggle) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// Slider is a simple slider with min and max values
type Slider struct {
	Name             *string `protobuf:"bytes,1,req,name=Name,json=name" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Slider) Reset()                    { *m = Slider{} }
func (m *Slider) String() string            { return proto.CompactTextString(m) }
func (*Slider) ProtoMessage()               {}
func (*Slider) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *Slider) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// Rotary is a rotaryencoder
type Rotary struct {
	Name             *string `protobuf:"bytes,1,req,name=Name,json=name" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rotary) Reset()                    { *m = Rotary{} }
func (m *Rotary) String() string            { return proto.CompactTextString(m) }
func (*Rotary) ProtoMessage()               {}
func (*Rotary) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *Rotary) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Component)(nil), "message.Component")
	proto.RegisterType((*Stat)(nil), "message.Stat")
	proto.RegisterType((*Toggle)(nil), "message.Toggle")
	proto.RegisterType((*Slider)(nil), "message.Slider")
	proto.RegisterType((*Rotary)(nil), "message.Rotary")
}

func init() { proto.RegisterFile("components.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0xda, 0xc5, 0xc8, 0xc5, 0xe9, 0x0c, 0x93, 0x15, 0x52, 0xe6,
	0x62, 0x09, 0x2e, 0x49, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd5, 0x83, 0xaa,
	0xd2, 0x03, 0x09, 0x7a, 0x30, 0x04, 0xb1, 0x14, 0x97, 0x24, 0x96, 0x08, 0x69, 0x72, 0xb1, 0x85,
	0xe4, 0xa7, 0xa7, 0xe7, 0xa4, 0x4a, 0x30, 0x81, 0x95, 0xf1, 0xc3, 0x95, 0x41, 0x84, 0x3d, 0x18,
	0x82, 0xd8, 0x4a, 0xc0, 0x2c, 0x90, 0xd2, 0xe0, 0x9c, 0xcc, 0x94, 0xd4, 0x22, 0x09, 0x66, 0x34,
	0xa5, 0x10, 0x61, 0x90, 0xd2, 0x62, 0x30, 0x0b, 0xa4, 0x34, 0x28, 0xbf, 0x24, 0xb1, 0xa8, 0x52,
	0x82, 0x05, 0x4d, 0x29, 0x44, 0x18, 0xa4, 0xb4, 0x08, 0xcc, 0x72, 0x62, 0xe7, 0x62, 0x2d, 0xcd,
	0xcb, 0xcc, 0xcf, 0x53, 0x92, 0x82, 0x38, 0x57, 0x48, 0x88, 0x8b, 0xc5, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x49, 0x83, 0x33, 0x88, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x49, 0x06, 0xe6, 0x4a,
	0x5c, 0xb2, 0x10, 0x17, 0xe0, 0x92, 0x85, 0x58, 0x8a, 0x4d, 0x16, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0xe6, 0xf7, 0x98, 0x4e, 0x01, 0x00, 0x00,
}
