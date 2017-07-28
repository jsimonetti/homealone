package main

import (
	"fmt"

	"golang.org/x/net/context"

	micro "github.com/micro/go-micro"
	natsbroker "github.com/micro/go-plugins/broker/nats"
	natsregistry "github.com/micro/go-plugins/registry/nats"
	natstransport "github.com/micro/go-plugins/transport/nats"

	"github.com/jsimonetti/homealone/services/messages"
)

func main() {
	registry := natsregistry.NewRegistry()
	brkr := natsbroker.NewBroker()
	trnsport := natstransport.NewTransport()

	service := micro.NewService(
		micro.Name("rpc"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(brkr),
		micro.Transport(trnsport),
	)
	service.Init()

	inventory := message.NewInventoryClient("inventory", service.Client())

	rsp, err := inventory.List(context.TODO(), &message.InventoryRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Printf("resp: %v\n", rsp)

	for key := range rsp.Devices {
		driver := message.NewDriverClient(key, service.Client())
		rsp2, err := driver.Execute(context.TODO(), &message.DriverCommand{
			Device: "test123",
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print response
		fmt.Printf("resp: %v\n", rsp2)

	}
}
