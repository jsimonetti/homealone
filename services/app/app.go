package app

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	natsbroker "github.com/micro/go-plugins/broker/nats"
	natsregistry "github.com/micro/go-plugins/registry/nats"
	natstransport "github.com/micro/go-plugins/transport/nats"
)

type App struct {
	name    string
	version string
	broker  broker.Broker
	service micro.Service
}

func NewApp(name, version string) *App {
	registry := natsregistry.NewRegistry()
	brkr := natsbroker.NewBroker()
	trnsport := natstransport.NewTransport()

	if err := brkr.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := brkr.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Registry(registry),
		micro.Broker(brkr),
		micro.Transport(trnsport),
		micro.RegisterInterval(30*time.Second),
		micro.RegisterTTL(65*time.Second),
	)
	service.Init()

	app := &App{
		name:    name,
		version: version,
		broker:  brkr,
		service: service,
	}

	return app
}

func (a *App) Service() micro.Service {
	return a.service
}

func (a *App) Broker() broker.Broker {
	return a.broker
}

func (a *App) Name() string {
	return a.name
}
func (a *App) Version() string {
	return a.version
}
