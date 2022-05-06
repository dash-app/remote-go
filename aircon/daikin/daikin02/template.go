package daikin02

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "daikin",
	Model:  "02",
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
			"auto": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 0.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   -5.0,
						To:     5.0,
						Suffix: "℃",
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto"},
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
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"cool": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 26.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   18.0,
						To:     32.0,
						Suffix: "℃",
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "1", "2", "3", "4", "5"},
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
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"dry": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 26.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   18.0,
						To:     32.0,
						Suffix: "℃",
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto"},
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
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"heat": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 23.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   14.0,
						To:     30.0,
						Suffix: "℃",
					},
				},
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "1", "2", "3", "4", "5"},
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
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
		},
	},
}

func (r *daikin02) Template() *appliances.Template {
	return Template
}
