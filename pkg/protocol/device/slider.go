package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Slider{}

// Slider is a simple slider with min and max values
type Slider struct {
	Name string
	Min  int
	Max  int
}

// component is an empty method to comply to the interface Component
func (Slider) component() {}

func init() {
	gob.Register(Slider{})
}
