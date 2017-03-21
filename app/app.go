// Package app contains helper functions to create application and drivers
package app

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
	"github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/logger"
	"github.com/jsimonetti/homealone/protocol"
	"github.com/jsimonetti/homealone/protocol/device"
	"github.com/jsimonetti/homealone/protocol/message"
	"github.com/jsimonetti/homealone/protocol/queue"
)

var host = flag.String("host", "localhost:1883", "hostname of broker")
var user = flag.String("user", "", "username")
var pass = flag.String("pass", "", "password")
var debug = flag.Bool("debug", false, "debugging enabled")

var namespace uuid.UUID

func init() {
	flag.Parse()

	// we hardcode the namespace here.
	namespace, _ = uuid.FromString("efdaedf2-4485-499f-8a68-3eea088fa7ae")
}

// App is a helper to create a driver or application.
// It takes away the burden of (un)registering devices at startup
// and responding to discovery messages.
type App struct {
	ID   uuid.UUID
	Log  log.Logger
	Name string

	debug bool

	conn   net.Conn
	Broker *mqtt.ClientConn

	Devices []device.Device

	shutdownCh chan struct{}
	Signal     chan os.Signal

	handler message.Handler
	wg      sync.WaitGroup
}

// NewApp returns an app in the 'app' namespace
func NewApp(name string) (*App, error) {
	return newApp("app." + name)
}

// NewDriver returns an app in the 'device' namespace
func NewDriver(name string) (*App, error) {
	return newApp("driver." + name)
}

// NewCore returns an app in the 'core' namespace
func NewCore(name string) (*App, error) {
	return newApp("core." + name)
}

// newApp returns an app
func newApp(name string) (app *App, err error) {
	app = &App{
		Name:    name,
		ID:      uuid.NewV5(namespace, "org.homealone."+name),
		handler: func(t string, m message.Message) error { return nil },
		debug:   *debug,
	}
	app.Log = log.NewLogger().With(log.Fields{"app": name, "id": app.ID})
	return
}

// Start starts this app. It connects to the message hub and registers devices.
// It also starts a goroutine to respond to discovery messages.
func (app *App) Start() (err error) {
	app.conn, err = net.Dial("tcp", *host)
	if err != nil {
		app.Log.WithError(err).Print("dial failed")
		return
	}

	app.Broker = mqtt.NewClientConn(app.conn)

	if err = app.Broker.Connect(*user, *pass); err != nil {
		app.Log.WithError(err).Print("connect failed")
		return
	}

	app.Log.With(log.Fields{"client_id": app.Broker.ClientId}).Print("client connected")

	app.RegisterAll()

	app.shutdownCh = make(chan struct{})
	go app.discoverLoop()

	app.Signal = make(chan os.Signal)
	signal.Notify(app.Signal, syscall.SIGINT, syscall.SIGTERM)

	return
}

// Stop will stop this app. It will stop the discovery goroutine,
// disconnect from the message hub and close its connection.
func (app *App) Stop() {
	app.UnregisterAll()
	close(app.shutdownCh)
	app.wg.Wait()
	app.Log.With(log.Fields{"client_id": app.Broker.ClientId}).Print("client disconnected")
	app.Broker.Disconnect()
	app.conn.Close()
}

// SetHandler will set the message handler
func (app *App) SetHandler(handler message.Handler) {
	app.handler = handler
}

// Publish will send a message to the specified topic.
func (app *App) Publish(topic queue.Topic, m message.Message) {
	b, err := protocol.Marshal(m)
	if err != nil {
		app.Log.WithError(err).Print("marshal failed")
	}

	app.Log.With(log.Fields{"topic": topic.String(), "source": m.From().String(), "type": m.Type().String()}).Print("sent message")
	app.Broker.Publish(&proto.Publish{
		Header:    proto.Header{},
		TopicName: topic.String(),
		Payload:   proto.BytesPayload(b),
	})
}

// DiscoverReply replies to a Discover message.
func (app *App) DiscoverReply() {
	m := &message.DiscoverReply{}
	m.Source = app.ID
	m.Name = app.Name
	m.Devices = app.Devices

	app.Publish(queue.Inventory, m)
}

// Wait will wait for an os.Signal to return
func (app *App) Wait() {
	<-app.Signal
}

// Debug will return the debug status of the app
func (app *App) Debug() bool {
	return app.debug
}

// Register will set the devices controlled by this app.
func (app *App) Register(device device.Device) {
	app.Devices = append(app.Devices, device)
}

// RegisterAll will send the Register message to the inventory.
func (app *App) RegisterAll() {
	m := &message.Register{}
	m.Source = app.ID
	m.Name = app.Name
	m.Devices = app.Devices

	app.Publish(queue.Inventory, m)
}

// UnregisterAll will unregister all devices from the inventory.
// This is usually done on shutdown.
func (app *App) UnregisterAll() {
	m := &message.Unregister{}
	m.Source = app.ID
	m.Name = app.Name
	m.Devices = app.Devices

	app.Publish(queue.Inventory, m)
}

// discoverLoop is the goroutine that replies to Discover messages
func (app *App) discoverLoop() {
	app.wg.Add(1)
	app.Broker.Subscribe([]proto.TopicQos{
		proto.TopicQos{
			Topic: queue.Inventory.String(),
			Qos:   proto.QosAtMostOnce,
		},
	})

	for {
		select {
		case <-app.shutdownCh:
			app.wg.Done()
			return

		case m := <-app.Broker.Incoming:
			msg, err := protocol.Unmarshal(m.Payload)
			if err != nil {
				app.Log.WithError(err).With(log.Fields{"topic": m.TopicName}).Print("unmarshal failed")
				break
			}
			app.Log.With(log.Fields{"topic": m.TopicName, "source": msg.From().String(), "type": msg.Type().String()}).Print("received message")
			if m.TopicName == queue.Inventory.String() {
				if msg.Type() == message.TypeDiscover {
					app.DiscoverReply()
					break
				}
			}
			if err := app.handler(m.TopicName, msg); err != nil {
				app.Log.WithError(err).Print("handler error")
			}
		}
	}
}
