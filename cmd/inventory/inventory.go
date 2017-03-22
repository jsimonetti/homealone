package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

func main() {

	app, err := app.NewCore("inventory")
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
	inventory := &InventoryApp{
		App:              app,
		discoverInterval: 30 * time.Second,
	}

	// decrease discovers when debugging
	if inventory.Debug() {
		inventory.discoverInterval = 5 * time.Second
	}

	inventory.SetHandler(inventory.messageHandler)
	inventory.Start()

	// wait for the interrupt signal
	inventory.Wait()

	// clean up
	close(inventory.shutdownCh)
	inventory.wg.Wait()
	inventory.Stop()
}

// InventoryApp is the inventory of devices and apps.
// It is the central yellowpages for the network.
// It will send periodic Discover messages to learn all apps and devices.
type InventoryApp struct {
	*app.App
	discoverInterval time.Duration

	shutdownCh chan struct{}
	wg         sync.WaitGroup
}

// Start will start the inventory app
func (app *InventoryApp) Start() {
	app.App.Start()
	m := &message.Discover{}
	m.Source = app.ID
	app.Publish(queue.Inventory, m)

	app.shutdownCh = make(chan struct{})
	go app.discoverLoop()
}

func (app *InventoryApp) discoverLoop() {
	app.wg.Add(1)
	timer := time.NewTicker(app.discoverInterval)

	m := &message.Discover{}
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

// inventoryReply will send the inventory of devices to the requester
func (app *InventoryApp) inventoryReply(to uuid.UUID) {
	m := &message.InventoryReply{}
	m.Source = app.ID
	m.For = to
	m.Devices = app.DeviceList()

	app.Publish(queue.Inventory, m)
}

func (app *InventoryApp) registerDevice(m *message.Register) {
	app.Register(m.Devices...)
}

func (app *InventoryApp) unregisterDevice(m *message.Unregister) {
	app.Unregister(m.Devices...)
}

// messageHandler is the handler to deal with messages
func (app *InventoryApp) messageHandler(topic string, m message.Message) error {
	switch m := m.(type) {
	case *message.Discover:
		return fmt.Errorf("unhandled message type %s", m.Type().String())
	case *message.Inventory:
		app.inventoryReply(m.From())
	case *message.Register:
		app.registerDevice(m)
	case *message.Unregister:
		app.unregisterDevice(m)
	default:
		fmt.Print(topic + "\t")
		spew.Dump(m)
	}
	return nil
}
