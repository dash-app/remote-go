package sg174

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "mitsubishi",
	Model:  "sg174",
	Kind:   appliances.AIRCON,
	Aircon: &appliances.AirconTemplate{
		Operation: &appliances.ActionTemplate{
			Type: appliances.TOGGLE,
			Toggle: &appliances.Toggle{
				ON:  true,
				OFF: false,
			},
		},
		Modes: map[string]*appliances.AirconModeTemplate{
			"cool": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 21.0,
					Range: &appliances.Range{
						Step: 0.5,
						From: 16.0,
						To:   31.0,
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "low", "mid", "high", "long"},
					},
				},
				HorizontalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"left", "left_mid", "center", "right_mid", "right"},
					},
				},
			},
			"dry": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "mid",
					List: &appliances.List{
						Values: []interface{}{"low", "mid", "high"},
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "low", "mid", "high", "long"},
					},
				},
				HorizontalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"left", "left_mid", "center", "right_mid", "right"},
					},
				},
			},
			"heat": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 27.0,
					Range: &appliances.Range{
						Step: 0.5,
						From: 16.0,
						To:   31.0,
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "low", "mid", "high", "long"},
					},
				},
				HorizontalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"left", "left_mid", "center", "right_mid", "right"},
					},
				},
			},
		},
	},
}

func (r *sg174) Template() *appliances.Template {
	return Template
}
