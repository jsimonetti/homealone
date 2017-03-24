package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jsimonetti/homealone/pkg/app"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

// Inventory is the application responsible for holding an inventory of all devices.
// It respons only to Discover, Inventory, Register and
// Unregister messages on the Inventory queue.

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

	inventory.SetHandler(queue.Inventory, inventory.messageHandler)
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
	m := &message.Discover{
		Header: &message.Header{
			From: &app.ID,
		},
	}
	app.Publish(queue.Inventory, m)

	app.shutdownCh = make(chan struct{})
	go app.discoverLoop()
}

// discoverLoop periodically asks for all devices and GC's stale devices
func (app *InventoryApp) discoverLoop() {
	app.wg.Add(1)
	timer := time.NewTicker(app.discoverInterval)

	m := &message.Discover{
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
			app.GCDeviceList(time.Now().Add(3 * app.discoverInterval))
		}
	}
}

// inventoryReply will send the inventory of devices to the requester
func (app *InventoryApp) inventoryReply(to string) error {
	m := &message.InventoryReply{
		Header: &message.Header{
			From: &app.ID,
			To:   &to,
		},
	}
	m.Devices = app.DeviceList()

	app.Publish(queue.Inventory, m)
	return nil
}

// registerDevice will register the devices to the inventory
func (app *InventoryApp) registerDevice(m *message.Register) error {
	app.Register(m.Devices...)
	return nil
}

// unregisterDevice will unregister the devices from the inventory
func (app *InventoryApp) unregisterDevice(m *message.Unregister) error {
	app.Unregister(m.Devices...)
	return nil
}

// messageHandler is the handler to deal with messages
func (app *InventoryApp) messageHandler(m message.Message) error {
	switch m := m.(type) {
	case *message.Discover:
		return fmt.Errorf("unhandled message type %s", m.Type().String())
	case *message.Inventory:
		return app.inventoryReply(m.From())
	case *message.Register:
		return app.registerDevice(m)
	case *message.Unregister:
		return app.unregisterDevice(m)
	}
	return nil
}
