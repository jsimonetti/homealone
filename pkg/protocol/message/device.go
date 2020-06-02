package message

// AddStat will add the named stat to the device
func (m *Device) AddStat(name string) {
	c := &Component{
		Union: &Component_Stat{
			Stat: &Stat{
				Name: &name,
			},
		},
	}
	m.Components = append(m.Components, c)
}

// AddSlider will add the named slider to the device
func (m *Device) AddSlider(name string) {
	c := &Component{
		Union: &Component_Slider{
			Slider: &Slider{
				Name: &name,
			},
		},
	}
	m.Components = append(m.Components, c)
}

// AddToggle will add the named toggle to the device
func (m *Device) AddToggle(name string) {
	c := &Component{
		Union: &Component_Toggle{
			Toggle: &Toggle{
				Name: &name,
			},
		},
	}
	m.Components = append(m.Components, c)
}

// AddRotary will add the named rotary to the device
func (m *Device) AddRotary(name string) {
	c := &Component{
		Union: &Component_Rotary{
			Rotary: &Rotary{
				Name: &name,
			},
		},
	}
	m.Components = append(m.Components, c)
}
