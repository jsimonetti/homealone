package message

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
