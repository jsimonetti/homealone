HomeAlone [![Build Status](https://travis-ci.org/jsimonetti/homealone.svg?branch=master)](https://travis-ci.org/jsimonetti/homealone) [![GoDoc](https://godoc.org/github.com/jsimonetti/homealone?status.svg)](https://godoc.org/github.com/jsimonetti/homealone) [![Go Report Card](https://goreportcard.com/badge/github.com/jsimonetti/homealone)](https://goreportcard.com/report/github.com/jsimonetti/homealone)
=======

HomeAlone is yet another home automation system written in golang.

It uses an MQTT broker for message communication between drivers.
An inventory keeps track of all devices in the network.
Each driver reports its' devices to the inventory and events to the event queue.
Commands can be sent via the command queue and each driver can react individually.

MIT Licensed.

This package is still unfinished. No working product yet.
I intend to at least add drivers for DMX/Art-Net and Zwave as these are the products I use.