package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// Drain is a helper program that prints out all messages on all queueus.
// It is for debugging purposes only

func main() {

	app, err := app.NewCore("drain")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	// drain listens on all topics and usage the dumpHandler to dump all messages
	for _, t := range queue.AllTopics() {
		app.SetHandler(t, dumpHandler)
	}
	app.FilterMessages(false)
	app.Start()

	// wait for the interrupt signal
	app.Wait()
}

// dumpHandler is the handler to deal with all messages
func dumpHandler(m message.Message) error {
	fmt.Printf("%s:  %s\n", time.Now(), m.String())
	return nil
}
