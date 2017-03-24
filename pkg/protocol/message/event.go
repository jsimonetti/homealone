package message

import uuid "github.com/satori/go.uuid"

// Used to make sure the interface is met
var _ Message = &Event{}

// Event is a message used by any app to signal an event
type Event struct {
	Header
	ID        uuid.UUID
	SubjectID uuid.UUID
	Event     EventType
	Message   string
}

// Finalize will finish the object before marshalling
func (m *Event) Finalize() {
	m.Header.MType = TypeEvent
}

// message is an empty method to comply to the interface Message
func (Event) message() {}

//go:generate stringer -type=EventType

// EventType is the type of the event.
type EventType uint8

// These constants define the different event types.
const (
	EventDriverStateChange EventType = iota
	EventDeviceStateChange
	EventComponentStateChange
	EventComponentValueChange
)
