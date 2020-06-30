package panasonic01

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "panasonic",
	Model:  "01",
	Kind:   "aircon",
	Aircon: &template.Aircon{
		Operation: &template.Action{
			Type: template.TOGGLE,
			Toggle: &template.Toggle{
				ON:  true,
				OFF: false,
			},
		},
		Modes: map[string]*template.AirconMode{
			"auto": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "left", "mid_left", "center", "mid_right", "right"},
				},
			},
			"cool": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "left", "mid_left", "center", "mid_right", "right"},
				},
			},
			"dry": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "left", "mid_left", "center", "mid_right", "right"},
				},
			},
			"heat": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 23.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "left", "mid_left", "center", "mid_right", "right"},
				},
			},
		},
	},
}

func (r *panasonic01) Template() *template.Template {
	return Template
}
