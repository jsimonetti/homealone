package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/davecgh/go-spew/spew"
	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"

	"github.com/jsimonetti/homealone/pkg/protocol"
	"github.com/jsimonetti/homealone/pkg/protocol/queue"
)

var host = flag.String("host", "localhost:1883", "hostname of broker")
var id = flag.String("id", "", "client id")
var user = flag.String("user", "", "username")
var pass = flag.String("pass", "", "password")
var dump = flag.Bool("dump", false, "dump messages?")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *host)
	if err != nil {
		print(err.Error() + "\n")
		return
	}
	cc := mqtt.NewClientConn(conn)
	cc.Dump = *dump
	cc.ClientId = *id

	if err := cc.Connect(*user, *pass); err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	var tq []proto.TopicQos

	for _, t := range queue.AllTopics() {
		tq = append(tq, proto.TopicQos{
			Topic: t.String(),
			Qos:   proto.QosAtMostOnce,
		})
	}
	cc.Subscribe(tq)

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		cc.Disconnect()
		conn.Close()
	}()

	for {
		select {
		case m := <-cc.Incoming:
			msg, err := protocol.Unmarshal(m.Payload)
			if err != nil {
				print(err.Error() + "\n")
				break
			}

			fmt.Print(m.TopicName, "\t")
			spew.Dump(msg)
		case <-osSignal:
			return
		}
	}
}
