package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Slider{}

// Slider is a simple slider with min and max values
type Slider struct {
	Name string
	min  int
	max  int
	val  int
}

// Set the value of i to the Value and return the new value.
// if value i is larger then max it is set to max, dito with min.
func (s *Slider) Set(i int) int {
	switch true {
	case i > s.max:
		s.val = s.max

	case i < s.min:
		s.val = s.min

	case s.max > i && i > s.min:
		s.val = i
	}
	return s.val
}

// Value returns the current value
func (s *Slider) Value() int {
	return s.val
}

// SliderCommands is an interface to what methods a Slider
// component should support
type SliderCommands interface {
	// Set the value of i to the Value and return the new value.
	// if value i is larger then max it is set to max, dito with min.
	Set(i int) int
	// Value returns the current value
	Value() int
}

// component is an empty method to comply to the interface Component
func (Slider) component() {}

func init() {
	gob.Register(Slider{})
}
