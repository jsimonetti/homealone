package device

import uuid "github.com/satori/go.uuid"

// Device holds information about a single device.
// A driver can supply multiple devices.
// UUID should be created deterministically in a way
// so that each startup the device gets the same uuid. (uuid.NewV5() helps here)
type Device struct {
	ID         uuid.UUID
	Name       string
	Components []Component
}

type Component interface {
	component()
}

type Toggle struct {
	Name string
}

func (Toggle) component() {}
