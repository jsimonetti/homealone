package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Rotary{}

// Rotary is a rotaryencoder
type Rotary struct {
	Name string
	wrap int
	val  int64
}

// Add the value of i to the Value and return the new value.
// The new value can wrap to an arbitrary number
func (r *Rotary) Add(i int) int {
	r.val += int64(i)

	if r.val > int64(r.wrap) {
		r.val = r.val - int64(r.wrap)
	}

	return int(r.val)
}

// Value returns the current value
func (r *Rotary) Value() int {
	return int(r.val)
}

// RotaryCommands is an interface to what methods a Rotary
// device should support
type RotaryCommands interface {
	// Add the value of i to the Value and return the new value.
	// The new value can wrap to an arbitrary number
	Add(i int) int
	// Value returns the current value
	Value() int
}

// component is an empty method to comply to the interface Component
func (Rotary) component() {}

func init() {
	gob.Register(Rotary{})
}
