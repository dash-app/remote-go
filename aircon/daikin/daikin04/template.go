package daikin04

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "daikin",
	Model:  "04",
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
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto"},
					},
				},
				HorizontalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
			"cool": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 26.0,
					Range: &appliances.Range{
						Step:   1.0,
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
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
			"dry": {
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto"},
					},
				},
				HorizontalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
			"heat": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 23.0,
					Range: &appliances.Range{
						Step:   1.0,
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
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
		},
	},
}

func (r *daikin04) Template() *appliances.Template {
	return Template
}
