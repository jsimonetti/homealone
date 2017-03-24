package protocol_test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/protocol"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
)

func TestProtocol(t *testing.T) {
	t.Run("Marshal", testProtocolMarshal)
	t.Run("Unmarshal", testProtocolUnmarshal)
}

func testProtocolMarshal(t *testing.T) {
	testUUID, _ := uuid.FromString("efdaedf2-4485-499f-8a68-3eea088fa7ae")
	tests := []struct {
		name string
		p    message.Message
		b    []byte
		err  error
	}{
		{
			name: "Empty Discover",
			p:    &message.Discover{},
			b: []byte{
				0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xff, 0x81, 0x03, 0x01, 0x01, 0x08,
				0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x01, 0xff, 0x82, 0x00, 0x01, 0x01, 0x01, 0x06,
				0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x00, 0x00, 0x33, 0xff, 0x83, 0x03,
				0x01, 0x01, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x01, 0x03, 0x01,
				0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x01, 0x06, 0x00, 0x01, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
				0x65, 0x01, 0xff, 0x86, 0x00, 0x01, 0x03, 0x46, 0x6f, 0x72, 0x01, 0xff, 0x86, 0x00, 0x00, 0x00,
				0x10, 0xff, 0x85, 0x06, 0x01, 0x01, 0x04, 0x55, 0x55, 0x49, 0x44, 0x01, 0xff, 0x86, 0x00, 0x00,
				0x00, 0x05, 0xff, 0x82, 0x01, 0x00, 0x00,
			},
		},
		{
			name: "Discover",
			p: &message.Discover{
				Header: message.Header{
					Source: testUUID,
					For:    testUUID,
				},
			},
			b: []byte{
				0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xff, 0x81, 0x03, 0x01, 0x01, 0x08,
				0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x01, 0xff, 0x82, 0x00, 0x01, 0x01, 0x01, 0x06,
				0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x00, 0x00, 0x33, 0xff, 0x83, 0x03,
				0x01, 0x01, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x01, 0x03, 0x01,
				0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x01, 0x06, 0x00, 0x01, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
				0x65, 0x01, 0xff, 0x86, 0x00, 0x01, 0x03, 0x46, 0x6f, 0x72, 0x01, 0xff, 0x86, 0x00, 0x00, 0x00,
				0x10, 0xff, 0x85, 0x06, 0x01, 0x01, 0x04, 0x55, 0x55, 0x49, 0x44, 0x01, 0xff, 0x86, 0x00, 0x00,
				0x00, 0x29, 0xff, 0x82, 0x01, 0x02, 0x10, 0xef, 0xda, 0xed, 0xf2, 0x44, 0x85, 0x49, 0x9f, 0x8a,
				0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7, 0xae, 0x01, 0x10, 0xef, 0xda, 0xed, 0xf2, 0x44, 0x85, 0x49,
				0x9f, 0x8a, 0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7, 0xae, 0x00, 0x00,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := protocol.Marshal(tt.p)

			if got, want := err, tt.err; got != want {
				t.Fatalf("unexpected error:\n - got: %v\n - want: %v\n", got, want)
			}
			if err != nil {
				return
			}

			if got, want := b, tt.b; !bytes.Equal(got, want) {
				t.Fatalf("unexpected Message bytes:\n - got:  [%s]\n - want: [%s]\n", printByteSlice(got), printByteSlice(want))
			}
		})
	}
}

func testProtocolUnmarshal(t *testing.T) {
	testUUID, _ := uuid.FromString("efdaedf2-4485-499f-8a68-3eea088fa7ae")
	tests := []struct {
		name string
		p    message.Message
		b    *payload
		err  error
	}{
		{
			name: "Version Mismatch",
			p:    &message.Discover{},
			b: &payload{
				[]byte{
					0x00, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xff, 0x81, 0x03, 0x01, 0x01, 0x08,
					0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x01, 0xff, 0x82, 0x00, 0x01, 0x01, 0x01, 0x06,
					0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x00, 0x00, 0x33, 0xff, 0x83, 0x03,
					0x01, 0x01, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x01, 0x03, 0x01,
					0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x01, 0x06, 0x00, 0x01, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
					0x65, 0x01, 0xff, 0x86, 0x00, 0x01, 0x03, 0x46, 0x6f, 0x72, 0x01, 0xff, 0x86, 0x00, 0x00, 0x00,
					0x10, 0xff, 0x85, 0x06, 0x01, 0x01, 0x04, 0x55, 0x55, 0x49, 0x44, 0x01, 0xff, 0x86, 0x00, 0x00,
					0x00, 0x05, 0xff, 0x82, 0x01, 0x00, 0x00,
				},
			},
			err: fmt.Errorf("protocol version mismatch"),
		},
		{
			name: "Empty Discover",
			p:    &message.Discover{},
			b: &payload{
				[]byte{
					0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xff, 0x81, 0x03, 0x01, 0x01, 0x08,
					0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x01, 0xff, 0x82, 0x00, 0x01, 0x01, 0x01, 0x06,
					0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x00, 0x00, 0x33, 0xff, 0x83, 0x03,
					0x01, 0x01, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x01, 0x03, 0x01,
					0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x01, 0x06, 0x00, 0x01, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
					0x65, 0x01, 0xff, 0x86, 0x00, 0x01, 0x03, 0x46, 0x6f, 0x72, 0x01, 0xff, 0x86, 0x00, 0x00, 0x00,
					0x10, 0xff, 0x85, 0x06, 0x01, 0x01, 0x04, 0x55, 0x55, 0x49, 0x44, 0x01, 0xff, 0x86, 0x00, 0x00,
					0x00, 0x05, 0xff, 0x82, 0x01, 0x00, 0x00,
				},
			},
		},
		{
			name: "Discover",
			p: &message.Discover{
				Header: message.Header{
					Source: testUUID,
					For:    testUUID,
				},
			},
			b: &payload{
				[]byte{
					0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xff, 0x81, 0x03, 0x01, 0x01, 0x08,
					0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x01, 0xff, 0x82, 0x00, 0x01, 0x01, 0x01, 0x06,
					0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x00, 0x00, 0x33, 0xff, 0x83, 0x03,
					0x01, 0x01, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x01, 0xff, 0x84, 0x00, 0x01, 0x03, 0x01,
					0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x01, 0x06, 0x00, 0x01, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
					0x65, 0x01, 0xff, 0x86, 0x00, 0x01, 0x03, 0x46, 0x6f, 0x72, 0x01, 0xff, 0x86, 0x00, 0x00, 0x00,
					0x10, 0xff, 0x85, 0x06, 0x01, 0x01, 0x04, 0x55, 0x55, 0x49, 0x44, 0x01, 0xff, 0x86, 0x00, 0x00,
					0x00, 0x29, 0xff, 0x82, 0x01, 0x02, 0x10, 0xef, 0xda, 0xed, 0xf2, 0x44, 0x85, 0x49, 0x9f, 0x8a,
					0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7, 0xae, 0x01, 0x10, 0xef, 0xda, 0xed, 0xf2, 0x44, 0x85, 0x49,
					0x9f, 0x8a, 0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7, 0xae, 0x00, 0x00,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := protocol.Unmarshal(tt.b)

			if got, want := err, tt.err; got != nil && want == nil {
				t.Fatalf("unexpected error:  %v\n", got)
			}
			if got, want := err, tt.err; got == nil && want != nil {
				t.Fatalf("expected error:  %v, got none", want)
			}
			if err != nil || tt.err != nil {
				return
			}

			tt.p.Finalize()

			if got, want := p, tt.p; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected Message:\n - got: [%#v]\n - want: [%#v]\n", got, want)
			}
		})
	}
}

var _ protocol.Payload = &payload{}

type payload struct{ b []byte }

func (p *payload) WritePayload(w io.Writer) error { _, err := w.Write(p.b); return err }

func printByteSlice(b []byte) string {
	str := ""
	i := 0
	for _, c := range b {
		if i%16 == 0 {
			str += "\n    "
		}
		str += fmt.Sprintf("%#02x, ", c)
		i++
	}
	str = "[]byte{" + str + "\n}"
	return str
}
