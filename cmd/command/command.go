package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
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

	driver, _ := uuid.FromString("ce0fc4f4-f62c-52c5-8747-e8c5035fdaed") // this is the dummy driver
	device, _ := uuid.FromString("42958300-8934-53a8-928f-e026531c2df8") // this is the device radio on dummy driver
	cmd := &message.Command{
		Header: message.Header{
			For:    driver,
			Source: app.ID,
		},
		ID:          uuid.NewV4(),
		Destination: device,
		Component:   "On/Off",
		Op:          "Toggle",
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
	fmt.Printf("%s\t", time.Now())
	spew.Dump(m)
	return nil
}

// sendCommand send the command
func (app *CommandApp) sendCommand(m message.Message) {
	app.Publish(queue.Command, m)
}
