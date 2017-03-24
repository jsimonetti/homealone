package main

import (
	"math/rand"
	"os"

	uuid "github.com/satori/go.uuid"

	"time"

	"github.com/jsimonetti/homealone/pkg/app"
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

func (app *DriverApp) fakeDevices() []*message.Device {
	baseuuid, _ := uuid.FromString(app.ID)
	name1 := "Lamp"
	id1 := uuid.NewV5(baseuuid, "Lamp").String()
	name2 := "Radio"
	id2 := uuid.NewV5(baseuuid, "Radio").String()

	devices := []*message.Device{
		&message.Device{
			ID:    &id1,
			Owner: &app.ID,
			Name:  &name1,
		},
		&message.Device{
			ID:    &id2,
			Owner: &app.ID,
			Name:  &name2,
		},
	}
	return devices
}

// commandHandler is the handler to deal with all messages
// from the command queue
func (app *DriverApp) commandHandler(m message.Message) error {
	switch m := m.(type) {
	case *message.Command:
		result, msg := app.executeDeviceOp(*m.Destination, *m.Op)
		reply := &message.CommandReply{
			Header: &message.Header{
				From: &app.ID,
				To:   m.Header.From,
			},
			InReplyTo: m.ID,
			Result:    &result,
			Message:   &msg,
		}
		app.Publish(queue.Command, reply)
		return nil
	}
	return nil
}

func (app *DriverApp) executeDeviceOp(id, op string) (message.CommandResult, string) {
	go app.sendEvent(id, message.EventType_ComponentValueChange, "change description")

	r := rand.Int() % 6
	return message.CommandResult(r), message.CommandResult(r).String()
}

func (app *DriverApp) sendEvent(id string, etype message.EventType, msg string) {
	time.Sleep(1 * time.Second)
	cmdid := uuid.NewV4().String()
	event := &message.Event{
		Header: &message.Header{
			From: &app.ID,
		},
		ID:        &cmdid,
		SubjectID: &id,
		Event:     &etype,
		Message:   &msg,
	}
	app.Publish(queue.Event, event)
}
