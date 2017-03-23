package main

import (
	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/protocol/device"
)

func fakeDevice() []device.Device {
	devices := []device.Device{
		device.Device{
			ID:   uuid.NewV5(namespace, "Lamp"),
			Name: "Lamp",
			Components: []device.Component{
				device.Toggle{
					Name: "On/Off",
				},
				device.Slider{
					Name: "Dimm",
				},
			},
		},
		device.Device{
			ID:   uuid.NewV5(namespace, "Radio"),
			Name: "Radio",
			Components: []device.Component{
				device.Toggle{
					Name: "On/Off",
				},
				device.Rotary{
					Name: "Volume",
				},
			},
		},
	}
	return devices
}
