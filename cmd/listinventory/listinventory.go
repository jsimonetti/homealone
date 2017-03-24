package main

import (
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// listinventory is a helper program that periodically prints out the current inventory.
// It is for debugging purposes only

func main() {

	app, err := app.NewCore("listinventory")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
	inventory := &ListInventoryApp{
		App:               app,
		inventoryInterval: 60 * time.Second,
	}

	// decrease discovers when debugging
	if inventory.Debug() {
		inventory.inventoryInterval = 10 * time.Second
	}

	inventory.SetHandler(queue.Inventory, inventoryHandler)
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
	m := &message.Inventory{
		Header: &message.Header{
			From: &app.ID,
		},
	}
	app.Publish(queue.Inventory, m)

	app.shutdownCh = make(chan struct{})
	go app.inventoryLoop()
}

func (app *ListInventoryApp) inventoryLoop() {
	app.wg.Add(1)
	timer := time.NewTicker(app.inventoryInterval)

	m := &message.Inventory{
		Header: &message.Header{
			From: &app.ID,
		},
	}

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

// inventoryHandler is the handler to deal with all messages
// except Discover (those are handled in the app framework)
func inventoryHandler(m message.Message) error {
	switch m := m.(type) {
	case *message.InventoryReply:
		spew.Dump(m)
	}
	return nil
}
