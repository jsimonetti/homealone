package main

import (
	"fmt"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// Command is a helper program that sends a command to the dummy driver
// It is for debugging purposes only

func main() {

	c, err := app.NewCore("command")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	app := &CommandApp{
		App: c,
	}

	app.SetHandler(queue.Command, commandHandler)
	app.Start()

	driver := "ce0fc4f4-f62c-52c5-8747-e8c5035fdaed" // this is the dummy driver
	//device, _ := uuid.FromString("42958300-8934-53a8-928f-e026531c2df8") // this is a non-existent device
	device := "615c4320-d2e2-5d8a-84b0-e3c62bc46feb" // this is the device radio on dummy driver

	comp := "Power"
	op := "Toggle"
	id := uuid.NewV4().String()

	cmd := &message.Command{
		Header: &message.Header{
			To:   &driver,
			From: &app.ID,
		},
		ID:          &id,
		Destination: &device,
		Component:   &comp,
		Op:          &op,
	}
	app.sendCommand(cmd)

	// wait for the interrupt signal
	app.Wait()
}

// CommandApp sends commands
type CommandApp struct {
	*app.App
}

// commandHandler is the handler to deal with all messages
func commandHandler(m message.Message) error {
	fmt.Printf("%s:  %s\n", time.Now(), m.String())
	return nil
}

// sendCommand send the command
func (app *CommandApp) sendCommand(m message.Message) {
	app.Publish(queue.Command, m)
}
