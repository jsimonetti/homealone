package protocol_test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/jsimonetti/homealone/pkg/protocol"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
)

func TestProtocol(t *testing.T) {
	t.Run("Marshal", testProtocolMarshal)
	t.Run("Unmarshal", testProtocolUnmarshal)
}

func testProtocolMarshal(t *testing.T) {
	from := "efdaedf2-4485-499f-8a68-3eea088fa7ae"
	to := "afdaedf2-4485-499f-8a68-3eea088fa7aa"
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
					From: &from,
					To:   &to,
				},
			},
			b: []byte{
				0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x01, 0x0a, 0x4e, 0x08, 0x01, 0x12, 0x24, 0x65,
				0x66, 0x64, 0x61, 0x65, 0x64, 0x66, 0x32, 0x2d, 0x34, 0x34, 0x38, 0x35, 0x2d, 0x34, 0x39, 0x39,
				0x66, 0x2d, 0x38, 0x61, 0x36, 0x38, 0x2d, 0x33, 0x65, 0x65, 0x61, 0x30, 0x38, 0x38, 0x66, 0x61,
				0x37, 0x61, 0x65, 0x1a, 0x24, 0x61, 0x66, 0x64, 0x61, 0x65, 0x64, 0x66, 0x32, 0x2d, 0x34, 0x34,
				0x38, 0x35, 0x2d, 0x34, 0x39, 0x39, 0x66, 0x2d, 0x38, 0x61, 0x36, 0x38, 0x2d, 0x33, 0x65, 0x65,
				0x61, 0x30, 0x38, 0x38, 0x66, 0x61, 0x37, 0x61, 0x61,
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
	from := "efdaedf2-4485-499f-8a68-3eea088fa7ae"
	to := "afdaedf2-4485-499f-8a68-3eea088fa7aa"
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
					Type: message.Type_discover.Enum(),
					From: &from,
					To:   &to,
				},
			},
			b: &payload{[]byte{
				0x68, 0x61, 0x2d, 0x31, 0x00, 0x00, 0x00, 0x00, 0x01, 0x0a, 0x4e, 0x08, 0x01, 0x12, 0x24, 0x65,
				0x66, 0x64, 0x61, 0x65, 0x64, 0x66, 0x32, 0x2d, 0x34, 0x34, 0x38, 0x35, 0x2d, 0x34, 0x39, 0x39,
				0x66, 0x2d, 0x38, 0x61, 0x36, 0x38, 0x2d, 0x33, 0x65, 0x65, 0x61, 0x30, 0x38, 0x38, 0x66, 0x61,
				0x37, 0x61, 0x65, 0x1a, 0x24, 0x61, 0x66, 0x64, 0x61, 0x65, 0x64, 0x66, 0x32, 0x2d, 0x34, 0x34,
				0x38, 0x35, 0x2d, 0x34, 0x39, 0x39, 0x66, 0x2d, 0x38, 0x61, 0x36, 0x38, 0x2d, 0x33, 0x65, 0x65,
				0x61, 0x30, 0x38, 0x38, 0x66, 0x61, 0x37, 0x61, 0x61,
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

			if got, want := p, tt.p; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected Message:\n - got:  [%s]\n - want: [%s]\n", got.String(), want.String())
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
