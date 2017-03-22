package main

import (
	"flag"
	"net"

	"github.com/jeffallen/mqtt"
)

var addr = flag.String("addr", "localhost:1883", "listen address of broker")

// Hub is the central MQTT broker to transport all messages.

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		print(err.Error() + "\n")
		return
	}
	svr := mqtt.NewServer(l)
	svr.Start()
	<-svr.Done
}
