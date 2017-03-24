package main

import (
	"os"

	uuid "github.com/satori/go.uuid"

	"time"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/device"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// The dummy driver is an example of a device driver.
// It is for debugging purposes only. It inserts some fake devices into the inventory.

func main() {

	c, err := app.NewDriver("dummy")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	app := &DriverApp{
		App: c,
	}

	app.Register(app.fakeDevices()...)
	app.SetHandler(queue.Command, app.commandHandler)
	app.Start()

	// wait for the interrupt signal
	app.Wait()

	app.Stop()
}

// DriverApp is this driver
type DriverApp struct {
	*app.App
}

func (app *DriverApp) fakeDevices() []device.Device {
	devices := []device.Device{
		device.Device{
			ID:    uuid.NewV5(app.ID, "Lamp"),
			Owner: app.ID,
			Name:  "Lamp",
			Components: []device.Component{
				device.Toggle{
					Name: "Power",
				},
				device.Slider{
					Name: "Dimm",
				},
			},
		},
		device.Device{
			ID:    uuid.NewV5(app.ID, "Radio"),
			Owner: app.ID,
			Name:  "Radio",
			Components: []device.Component{
				device.Toggle{
					Name: "Power",
				},
				device.Rotary{
					Name: "Volume",
				},
			},
		},
	}
	return devices
}

// commandHandler is the handler to deal with all messages
// from the command queue
func (app *DriverApp) commandHandler(m message.Message) error {
	switch m := m.(type) {
	case *message.Command:
		result, msg := app.executeDeviceOp(m.Destination, m.Op)
		reply := &message.CommandReply{
			Header: message.Header{
				Source: app.ID,
				For:    m.Source,
			},
			InReplyTo: m.ID,
			Result:    result,
			Message:   msg,
		}
		app.Publish(queue.Command, reply)
		return nil
	}
	return nil
}

func (app *DriverApp) executeDeviceOp(id uuid.UUID, op string) (message.CommandResult, string) {
	go app.sendEvent(id, message.EventComponentValueChange, "change description")
	return message.CommandSyncAck, ""
}

func (app *DriverApp) sendEvent(id uuid.UUID, etype message.EventType, msg string) {
	time.Sleep(1 * time.Second)
	event := &message.Event{
		Header: message.Header{
			Source: app.ID,
		},
		ID:        uuid.NewV4(),
		SubjectID: id,
		Event:     etype,
		Message:   msg,
	}
	app.Publish(queue.Event, event)
}
