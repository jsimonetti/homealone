package main

import (
	"os"

	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/app"
	"github.com/jsimonetti/homealone/protocol/device"
	"github.com/jsimonetti/homealone/protocol/message"
)

func main() {

	app, err := app.NewDriver("dummy")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	app.SetHandler(messageHandler)
	app.Register(fakeDevice())
	app.Start()

	// wait for the interrupt signal
	app.Wait()

	app.Stop()
}

// messageHandler is the handler to deal with all messages
// except Discover (those are handled in the app framework)
func messageHandler(topic string, m message.Message) error {
	switch m.Type() {
	default:
		return nil
	}
}

func fakeDevice() device.Device {
	d := device.Device{
		ID:   uuid.NewV4(),
		Name: "Lamp",
		Components: []device.Component{
			device.Toggle{
				Name: "On/Off",
			},
		},
	}
	return d
}
