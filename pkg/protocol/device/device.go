package device

import uuid "github.com/satori/go.uuid"

// Device holds information about a single device.
// A driver can supply multiple devices.
// UUID should be created deterministically in a way
// so that each startup the device gets the same uuid. (uuid.NewV5() helps here)
type Device struct {
	ID         uuid.UUID
	Owner      uuid.UUID
	Name       string
	Components []Component
}

// Component is an interface to device components
type Component interface {
	component()
}
