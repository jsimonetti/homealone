package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/davecgh/go-spew/spew"
	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"

	"github.com/jsimonetti/homealone/logger"
	"github.com/jsimonetti/homealone/protocol"
	"github.com/jsimonetti/homealone/protocol/queue"
)

var host = flag.String("host", "localhost:1883", "hostname of broker")
var id = flag.String("id", "", "client id")
var user = flag.String("user", "", "username")
var pass = flag.String("pass", "", "password")
var dump = flag.Bool("dump", false, "dump messages?")

var logger log.Logger

func main() {
	flag.Parse()
	logger = log.NewLogger().With(log.Fields{"app": "drain"})

	conn, err := net.Dial("tcp", *host)
	if err != nil {
		logger.WithError(err).Print("dial failed")
		return
	}
	cc := mqtt.NewClientConn(conn)
	cc.Dump = *dump
	cc.ClientId = *id

	if err := cc.Connect(*user, *pass); err != nil {
		logger.WithError(err).Print("connect failed")
		os.Exit(1)
	}

	logger.With(log.Fields{"client_id": cc.ClientId}).Print("client connected")

	var tq []proto.TopicQos

	for _, t := range queue.AllTopics() {
		tq = append(tq, proto.TopicQos{
			Topic: t.String(),
			Qos:   proto.QosAtMostOnce,
		})
	}
	cc.Subscribe(tq)

	for m := range cc.Incoming {
		msg, err := protocol.Unmarshal(m.Payload)
		if err != nil {
			logger.WithError(err).With(log.Fields{"topic": m.TopicName}).Print("unmarshal failed")
			continue
		}

		fmt.Print(m.TopicName, "\t")
		spew.Dump(msg)
	}
}
