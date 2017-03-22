package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

func main() {

	app, err := app.NewCore("listinventory")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
	inventory := &ListInventoryApp{
		App:               app,
		inventoryInterval: 30 * time.Second,
	}

	// decrease discovers when debugging
	if inventory.Debug() {
		inventory.inventoryInterval = 5 * time.Second
	}

	inventory.SetHandler(messageHandler)
	inventory.Start()

	// wait for the interrupt signal
	inventory.Wait()

	// clean up
	close(inventory.shutdownCh)
	inventory.wg.Wait()
	inventory.Stop()
}

// ListInventoryApp periodically lists the inventory
type ListInventoryApp struct {
	*app.App
	inventoryInterval time.Duration

	shutdownCh chan struct{}
	wg         sync.WaitGroup
}

// Start will start the inventory app
func (app *ListInventoryApp) Start() {
	app.App.Start()
	m := &message.Inventory{}
	m.Source = app.ID
	app.Publish(queue.Inventory, m)

	app.shutdownCh = make(chan struct{})
	go app.inventoryLoop()
}

func (app *ListInventoryApp) inventoryLoop() {
	app.wg.Add(1)
	timer := time.NewTicker(app.inventoryInterval)

	m := &message.Inventory{}
	m.Source = app.ID

	for {
		select {
		case <-app.shutdownCh:
			app.wg.Done()
			return

		case <-timer.C:
			app.Publish(queue.Inventory, m)
		}
	}
}

// messageHandler is the handler to deal with all messages
// except Discover (those are handled in the app framework)
func messageHandler(topic string, m message.Message) error {
	switch m := m.(type) {
	case *message.InventoryReply:
		fmt.Print(topic + "\t")
		spew.Dump(m)
	}
	return nil
}
