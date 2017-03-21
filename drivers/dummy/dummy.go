package main

import (
	"os"

	"github.com/jsimonetti/homealone/app"
	"github.com/jsimonetti/homealone/protocol/message"
)

func main() {

	app, err := app.NewDriver("dummy")
	if err != nil {
		os.Exit(1)
	}

	app.SetHandler(messageHandler)
	app.Start()

	// wait for the interupt signal
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
