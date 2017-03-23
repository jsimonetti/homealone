package main

import (
	"os"

	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/device"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// The dummy driver is an example of a device driver.
// It is for debugging purposes only. It inserts some fake devices into the inventory.

// namespace contains a UUID for this driver
var namespace uuid.UUID

func init() {
	namespace, _ = uuid.FromString("ffdaedf2-4485-499f-8a68-3eea088fa7ae")
}

func main() {

	c, err := app.NewDriver("dummy")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	app := &DriverApp{
		App: c,
	}

	app.Register(fakeDevice()...)
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

// commandHandler is the handler to deal with all messages
// from the command queue
func (app *DriverApp) commandHandler(m message.Message) error {
	switch m := m.(type) {
	case *message.Command:
		// filter to only react on my devices
		for _, device := range app.DeviceList() {
			if uuid.Equal(m.Destination, device.ID) {
				result, msg := app.execute(device, m.Op)
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
		}
		// if this was unicasted to us, we should reply
		if uuid.Equal(m.For, app.ID) {
			reply := &message.CommandReply{
				Header: message.Header{
					Source: app.ID,
					For:    m.Source,
				},
				InReplyTo: m.ID,
				Result:    message.CommandSyncFail,
				Message:   "no such device found",
			}
			app.Publish(queue.Command, reply)
		}
	}
	return nil
}

func (app *DriverApp) execute(device device.Device, op string) (message.CommandResult, string) {
	return message.CommandSyncAck, ""
}
