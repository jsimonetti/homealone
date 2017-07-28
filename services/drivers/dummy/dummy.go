package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/davecgh/go-spew/spew"

	"github.com/jsimonetti/homealone/services/app"
	"github.com/jsimonetti/homealone/services/messages"
	"github.com/micro/go-micro/broker"
)

// Setup and the client
func main() {
	driver := &Driver{}

	driver.App = app.NewApp("driver.dummy", "latest")
	message.RegisterDriverHandler(driver.Service().Server(), driver)

	driver.sub()

	if err := driver.Service().Run(); err != nil {
		log.Fatal(err)
	}
}

func (g *Driver) sub() {
	_, err := g.Broker().Subscribe("go.micro.topic.inventory", func(p broker.Publication) error {
		if string(p.Message().Body) == "Discover" {
			inventory := message.NewInventoryClient(p.Message().Header["source"], g.Service().Client())
			register(inventory)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

type Driver struct {
	*app.App

	inventory message.InventoryClient

	devices []*message.Device
}

func (g *Driver) Execute(ctx context.Context, req *message.DriverCommand, rsp *message.CommandResult) error {
	spew.Dump(req)
	rsp.Result = "OK"
	return nil
}

func register(inventory message.InventoryClient) {
	resp, err := inventory.Register(context.TODO(), &message.RegisterRequest{
		Devices: []*message.Device{
			&message.Device{
				ID:   "12345",
				Name: "Lamp",
				Components: []*message.Component{
					&message.Component{
						Union: &message.Component_Rotary{
							Rotary: &message.Rotary{
								Name: "Volume",
							},
						},
					},
					&message.Component{
						Union: &message.Component_Toggle{
							Toggle: &message.Toggle{
								Name: "Power",
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
