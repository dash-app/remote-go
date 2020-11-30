package sg174

import (
	"github.com/dash-app/remote-go/template"
)

var Template = &template.Template{
	Vendor: "mitsubishi",
	Model:  "sg174",
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
			"cool": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 21.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   31.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "low", "mid", "high", "long"},
					},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "center",
					List: &template.List{
						Values: []interface{}{"left", "left_mid", "center", "right_mid", "right"},
					},
				},
			},
			"dry": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.LIST,
					Default: "mid",
					List: &template.List{
						Values: []interface{}{"low", "mid", "high"},
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "low", "mid", "high", "long"},
					},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "center",
					List: &template.List{
						Values: []interface{}{"left", "left_mid", "center", "right_mid", "right"},
					},
				},
			},
			"heat": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 27.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   31.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "low", "mid", "high", "long"},
					},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "center",
					List: &template.List{
						Values: []interface{}{"left", "left_mid", "center", "right_mid", "right"},
					},
				},
			},
		},
	},
}

func (r *sg174) Template() *template.Template {
	return Template
}
