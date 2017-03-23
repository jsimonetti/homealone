package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/device"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// The dummy driver is an example of a device driver.
// It is for debugging purposes only. It inserts some fake devices into the inventory.

func main() {

	app, err := app.NewDriver("dummy")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	app.SetHandler(queue.Command, commandHandler)
	app.Register(fakeDevice()...)
	app.Start()

	// wait for the interrupt signal
	app.Wait()

	app.Stop()
}

// commandHandler is the handler to deal with all messages
// from the command queue
func commandHandler(m message.Message) error {
	switch m.Type() {
	default:
		spew.Dump(m)
		return nil
	}
}

func fakeDevice() []device.Device {
	devices := []device.Device{
		device.Device{
			ID:   uuid.NewV4(),
			Name: "Lamp",
			Components: []device.Component{
				device.Toggle{
					Name: "On/Off",
				},
				device.Slider{
					Name: "Dimm",
				},
			},
		},
		device.Device{
			ID:   uuid.NewV4(),
			Name: "Radio",
			Components: []device.Component{
				device.Toggle{
					Name: "On/Off",
				},
				device.Rotary{
					Name: "Volume",
				},
			},
		},
	}
	return devices
}
