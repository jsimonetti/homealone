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
			name: "Discover",
			p: &message.Discover{
				Header: &message.Header{
					Mfrom: testUUID.Bytes(),
					Mto:   testUUID.Bytes(),
				},
			},
			b: []byte{
				0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x01, 0x0a, 0x26, 0x08, 0x01, 0x12, 0x10, 0xef,
				0xda, 0xed, 0xf2, 0x44, 0x85, 0x49, 0x9f, 0x8a, 0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7, 0xae, 0x1a,
				0x10, 0xef, 0xda, 0xed, 0xf2, 0x44, 0x85, 0x49, 0x9f, 0x8a, 0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7,
				0xae,
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
					0x00, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00,
				},
			},
			err: fmt.Errorf("protocol version mismatch"),
		},
		{
			name: "Discover",
			p: &message.Discover{
				Header: &message.Header{
					Mfrom: testUUID.Bytes(),
					Mto:   testUUID.Bytes(),
				},
			},
			b: &payload{
				[]byte{
					0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x01, 0x0a, 0x26, 0x08, 0x01, 0x12, 0x10, 0xef,
					0xda, 0xed, 0xf2, 0x44, 0x85, 0x49, 0x9f, 0x8a, 0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7, 0xae, 0x1a,
					0x10, 0xef, 0xda, 0xed, 0xf2, 0x44, 0x85, 0x49, 0x9f, 0x8a, 0x68, 0x3e, 0xea, 0x08, 0x8f, 0xa7,
					0xae,
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
