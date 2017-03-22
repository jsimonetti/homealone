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
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/logger"
	"github.com/jsimonetti/homealone/pkg/protocol"
	"github.com/jsimonetti/homealone/pkg/protocol/device"
	"github.com/jsimonetti/homealone/pkg/protocol/message"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
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

	devices    []device.Device
	deviceLock sync.RWMutex

	shutdownCh chan struct{}
	Signal     chan os.Signal

	handler message.Handler
	wg      sync.WaitGroup
}

// NewApp returns an app in the 'app' namespace
func NewApp(name string) (*App, error) {
	return newApp("app." + name)
}

// NewDriver returns an app in the 'driver' namespace
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
	return app, errors.Wrap(err, "newApp failed")
}

// Start starts this app. It connects to the message hub and registers devices.
// It also starts a goroutine to respond to discovery messages.
func (app *App) Start() (err error) {
	app.conn, err = net.Dial("tcp", *host)
	if err != nil {
		app.Log.WithError(err).Print("dial failed")
		return errors.Wrap(err, "dial failed")
	}

	app.Broker = mqtt.NewClientConn(app.conn)

	if err = app.Broker.Connect(*user, *pass); err != nil {
		app.Log.WithError(err).Print("connect failed")
		return errors.Wrap(err, "connect failed")
	}

	app.Log.With(log.Fields{"client_id": app.Broker.ClientId}).Print("client connected")

	app.RegisterAll(uuid.Nil)

	app.shutdownCh = make(chan struct{})
	go app.discoverLoop()

	app.Signal = make(chan os.Signal, 1)
	signal.Notify(app.Signal, syscall.SIGINT, syscall.SIGTERM)

	return errors.Wrap(err, "App start failed")
}

// Stop will stop this app. It will stop the discovery goroutine,
// disconnect from the message hub and close its connection.
func (app *App) Stop() {
	app.UnregisterAll(uuid.Nil)
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
		return
	}

	app.Log.With(log.Fields{"topic": topic.String(), "destination": m.To().String(), "source": m.From().String(), "type": m.Type().String()}).Print("sent message")

	app.Broker.Publish(&proto.Publish{
		Header:    proto.Header{},
		TopicName: topic.String(),
		Payload:   proto.BytesPayload(b),
	})
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
func (app *App) Register(devices ...device.Device) {
	app.deviceLock.Lock()
	defer app.deviceLock.Unlock()
	for _, device := range devices {
		found := false
		for _, d := range app.devices {
			if device.ID == d.ID {
				found = true
			}
		}
		if !found {
			app.devices = append(app.devices, device)
		}
	}
}

// Unregister will remove the devices from this app.
func (app *App) Unregister(devices ...device.Device) {
	app.deviceLock.Lock()
	defer app.deviceLock.Unlock()
	for _, device := range devices {
		for i, d := range app.devices {
			if device.ID == d.ID {
				app.devices = append(app.devices[:i], app.devices[i+1:]...)
			}
		}
	}
}

// RegisterAll will send the Register message to the inventory.
func (app *App) RegisterAll(to uuid.UUID) {
	m := &message.Register{}
	m.Source = app.ID
	m.Name = app.Name
	m.For = to
	m.Devices = app.DeviceList()

	app.Publish(queue.Inventory, m)
}

// UnregisterAll will unregister all devices from the inventory.
// This is usually done on shutdown.
func (app *App) UnregisterAll(to uuid.UUID) {
	m := &message.Unregister{}
	m.Source = app.ID
	m.Name = app.Name
	m.For = to
	m.Devices = app.DeviceList()

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
			app.Log.With(log.Fields{"topic": m.TopicName, "destination": msg.To().String(), "source": msg.From().String(), "type": msg.Type().String()}).Print("received message")

			// only reply to broadcasts or msgs directed to me
			if uuid.Equal(msg.To(), app.ID) || msg.To() == uuid.Nil {
				if m.TopicName == queue.Inventory.String() {
					if msg.Type() == message.TypeDiscover {
						app.RegisterAll(msg.From())
						break
					}
				}
				if err := app.handler(m.TopicName, msg); err != nil {
					app.Log.WithError(err).Print("handler error")
				}
			}
		}
	}
}

// DeviceList will return the list of devices
func (app *App) DeviceList() []device.Device {
	app.deviceLock.RLock()
	defer app.deviceLock.RUnlock()
	return app.devices
}
