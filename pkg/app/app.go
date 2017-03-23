package app

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

type appDevice struct {
	device   device.Device
	lastSeen time.Time
}

// App is a helper to create a driver or application.
// It takes away the burden of (un)registering devices at startup
// and responding to discovery messages.
type App struct {
	ID   uuid.UUID
	Log  log.Logger
	Name string

	debug          bool
	filterMessages bool

	conn   net.Conn
	Broker *mqtt.ClientConn

	devices    []appDevice
	deviceLock sync.RWMutex

	shutdownCh chan struct{}
	Signal     chan os.Signal

	handler map[queue.Topic]message.Handler
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
		Name:           name,
		ID:             uuid.NewV5(namespace, "org.homealone."+name),
		handler:        make(map[queue.Topic]message.Handler),
		debug:          *debug,
		filterMessages: true,
	}
	app.Log = log.NewLogger().With(log.Fields{"app": name, "id": app.ID})
	return app, errors.Wrap(err, "newApp failed")
}

// FilterMessages allows the app to control automatic message filtering
// By default messages are filtered
func (app *App) FilterMessages(b bool) {
	app.filterMessages = b
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

	app.Log.With(log.Fields{"client_id": app.Broker.ClientId}).Print("client started")

	app.RegisterAll(uuid.Nil)

	app.shutdownCh = make(chan struct{})
	go app.messageLoop()

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
	app.Log.With(log.Fields{"client_id": app.Broker.ClientId}).Print("client stopped")
	app.Broker.Disconnect()
	app.conn.Close()
}

// SetHandler will set the message handler
func (app *App) SetHandler(topic queue.Topic, handler message.Handler) {
	app.handler[topic] = handler
}

// Publish will send a message to the specified topic.
func (app *App) Publish(topic queue.Topic, m message.Message) {
	b, err := protocol.Marshal(m)
	if err != nil {
		app.Log.WithError(err).Print("marshal failed")
		return
	}

	if app.debug {
		app.Log.With(log.Fields{"topic": topic.String(), "destination": m.To().String(), "source": m.From().String(), "type": m.Type().String()}).Print("sent message")
	}

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
	// we are registering devices, we should also register for
	// discovery messages, but they are handled in the app.
	// we add a NoopHandler to the map to make the subscription work
	// we only do this if no handler is allready registered
	if _, ok := app.handler[queue.Inventory]; !ok {
		app.SetHandler(queue.Inventory, app.discoveryHandler)
	}

	app.deviceLock.Lock()
	defer app.deviceLock.Unlock()
	for _, device := range devices {
		found := false
		for _, d := range app.devices {
			if device.ID == d.device.ID {
				d.lastSeen = time.Now()
				found = true
			}
		}
		if !found {
			app.devices = append(app.devices, appDevice{device: device, lastSeen: time.Now()})
		}
	}
}

// Unregister will remove the devices from this app.
func (app *App) Unregister(devices ...device.Device) {
	app.deviceLock.Lock()
	defer app.deviceLock.Unlock()
	for _, device := range devices {
		for i, d := range app.devices {
			if device.ID == d.device.ID {
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

// messageLoop is the goroutine that receives messages
func (app *App) messageLoop() {
	app.wg.Add(1)

	// only subsribe to topics we have handlers for
	var topics []proto.TopicQos
	for topic, _ := range app.handler {
		topics = append(topics, proto.TopicQos{
			Topic: topic.String(),
			Qos:   proto.QosAtMostOnce,
		})
	}
	app.Broker.Subscribe(topics)

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

			if app.debug {
				app.Log.With(log.Fields{"topic": m.TopicName, "destination": msg.To().String(), "source": msg.From().String(), "type": msg.Type().String()}).Print("received message")
			}

			topic := queue.GetTopic(m.TopicName)

			// only reply to broadcasts or msgs directed to me
			if !app.filterMessages || msg.To() == uuid.Nil || uuid.Equal(msg.To(), app.ID) {
				// get handler from the handler map
				if handle, ok := app.handler[topic]; ok {
					if err := handle(msg); err != nil {
						app.Log.WithError(err).Print("handler error")
					}
					break
				}
				fmt.Printf("no handler found\n")
			}
		}
	}
}

// DeviceList will return the list of devices
func (app *App) DeviceList() []device.Device {
	app.deviceLock.RLock()
	defer app.deviceLock.RUnlock()
	devices := []device.Device{}
	for _, d := range app.devices {
		devices = append(devices, d.device)
	}
	return devices
}

// GCDeviceList will remove all devices older then time t
func (app *App) GCDeviceList(t time.Time) {
	app.deviceLock.Lock()
	defer app.deviceLock.Unlock()
start:
	for i, d := range app.devices {
		if d.lastSeen.Before(t) {
			app.devices = append(app.devices[:i], app.devices[i+1:]...)
			goto start
		}
	}
}

// NoopHandler is a function that noops for a specific message
func NoopHandler(m message.Message) error {
	return nil
}

// discoveryHandler is a function that noops for a specific message
func (app *App) discoveryHandler(m message.Message) error {
	app.RegisterAll(m.From())
	return nil
}
