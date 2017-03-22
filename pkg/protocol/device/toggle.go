package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Toggle{}

// Toggle is a simple binary toggle
type Toggle struct {
	Name string
	val  bool
}

// Toggle will toggle the component and returns the current value
func (t *Toggle) Toggle() bool {
	t.val = !t.val
	return t.val
}

// On turns the component on and returns the current value
func (t *Toggle) On() bool {
	t.val = true
	return true
}

// Off turns the component off and returns the current value
func (t *Toggle) Off() bool {
	t.val = false
	return false
}

// Value returns the current value
func (t *Toggle) Value() bool {
	return t.val
}

// ToggleCommands is an interface to what methods a Toggle
// component should support
type ToggleCommands interface {
	// Toggle will toggle the component and returns the current value
	Toggle() bool
	// On turns the component on and returns the current value
	On() bool
	// Off turns the component off and returns the current value
	Off() bool
	// Value returns the current value
	Value() bool
}

// component is an empty method to comply to the interface Component
func (Toggle) component() {}

func init() {
	gob.Register(Toggle{})
}
