package main

import (
	uuid "github.com/satori/go.uuid"

	"github.com/jsimonetti/homealone/pkg/protocol/device"
)

func (app *DriverApp) fakeDevice() []device.Device {
	devices := []device.Device{
		device.Device{
			ID:   uuid.NewV5(app.ID, "Lamp"),
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
			ID:   uuid.NewV5(app.ID, "Radio"),
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
