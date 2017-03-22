package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Rotary{}

// Rotary is a rotaryencoder
type Rotary struct {
	Name string
}

// component is an empty method to comply to the interface Component
func (Rotary) component() {}

func init() {
	gob.Register(Rotary{})
}
