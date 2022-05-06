package mitsubishi02

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "mitsubishi",
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
			"cool": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 26.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   16.0,
						To:     31.0,
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
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"dry": {
				Humid: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "50%",
					List: &appliances.List{
						Values: []interface{}{"40%", "50%", "60%"},
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
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"heat": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 26.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   16.0,
						To:     31.0,
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
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"fan": {
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
					Default: "center",
					List: &appliances.List{
						Values: []interface{}{"swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
		},
	},
}

func (r *mitsubishi02) Template() *appliances.Template {
	return Template
}
