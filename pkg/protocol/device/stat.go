package device

import "encoding/gob"

// Used to make sure the interface is met
var _ Component = &Stat{}

// Stat is a statistic and can be any value you desire
type Stat struct {
	Name string
	val  interface{}
}

// Value returns the current value
func (s *Stat) Value() interface{} {
	return s.val
}

// StatCommands is an interface to what methods a Stat
// componen should support
type StatCommands interface {
	// Value returns the current value
	Value() int
}

// component is an empty method to comply to the interface Component
func (Stat) component() {}

func init() {
	gob.Register(Stat{})
}
