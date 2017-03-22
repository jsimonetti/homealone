package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Toggle{}

// Toggle is a simple binary toggle
type Toggle struct {
	Name string
}

// component is an empty method to comply to the interface Component
func (Toggle) component() {}

func init() {
	gob.Register(Toggle{})
}
