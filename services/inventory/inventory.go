package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/jsimonetti/homealone/services/app"
	"github.com/jsimonetti/homealone/services/messages"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/metadata"
)

func main() {

	inventory := &Inventory{}

	inventory.App = app.NewApp("inventory", "latest")
	inventory.discoverInterval = 30 * time.Second

	message.RegisterInventoryHandler(inventory.Service().Server(), inventory)

	// Start
	inventory.shutdownCh = make(chan struct{})
	go inventory.discoverLoop()

	if err := inventory.Service().Run(); err != nil {
		log.Fatal(err)
	}

	// Shutdown
	close(inventory.shutdownCh)
	inventory.wg.Wait()
}

type Inventory struct {
	*app.App
	discoverInterval time.Duration
	shutdownCh       chan struct{}
	wg               sync.WaitGroup

	inventoryLock sync.RWMutex
	inventory     map[string][]*message.Device
}

// discoverLoop periodically asks for all devices and GC's stale devices
func (app *Inventory) discoverLoop() {
	app.wg.Add(1)
	timer := time.NewTicker(app.discoverInterval)

	msg := &broker.Message{
		Header: map[string]string{
			"source": app.Name(),
		},
		Body: []byte("Discover"),
	}

	go func() {
		<-time.After(5 * time.Second)
		if err := app.Broker().Publish("go.micro.topic.inventory", msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		}
	}()

	for {
		select {
		case <-app.shutdownCh:
			app.wg.Done()
			return

		case <-timer.C:
			if err := app.Broker().Publish("go.micro.topic.inventory", msg); err != nil {
				log.Printf("[pub] failed: %v", err)
			}
		}
	}
}

func (g *Inventory) List(ctx context.Context, req *message.InventoryRequest, rsp *message.InventoryResponse) error {
	rsp.Devices = make(map[string]*message.Devices)
	g.inventoryLock.RLock()
	for k := range g.inventory {
		rsp.Devices[k] = &message.Devices{
			Devices: g.inventory[k],
		}
	}
	g.inventoryLock.RUnlock()
	return nil
}

func (g *Inventory) Register(ctx context.Context, req *message.RegisterRequest, rsp *message.RegisterResponse) error {
	var source string
	var ok bool

	md, _ := metadata.FromContext(ctx)
	if source, ok = md["X-Micro-From-Service"]; !ok {
		return fmt.Errorf("no source in ctx")
	}

	g.inventoryLock.Lock()
	if g.inventory[source] == nil {
		g.inventory = make(map[string][]*message.Device)
	}
	for _, device := range req.Devices {
		found := false
		for _, d := range g.inventory[source] {
			if device.ID == d.ID {
				found = true
				break
			}
		}
		if !found {
			g.inventory[source] = append(g.inventory[source], device)
		}
	}
	g.inventoryLock.Unlock()
	return nil
}
