package sc4ua

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "mitsubishi",
	Model:  "PAR-SC4UA",
	Kind:   appliances.AIRCON,
	Aircon: &appliances.AirconTemplate{
		Operation: &appliances.ActionTemplate{
			Type: appliances.TOGGLE,
			Toggle: &appliances.Toggle{
				ON:  true,
				OFF: false,
			},
		},
		Options: map[string]*appliances.ActionTemplate{
			"econo": {
				Type:    appliances.LIST,
				Default: "OFF",
				List: &appliances.List{
					Values: []interface{}{"OFF", "ON"},
				},
			},
		},
		Modes: map[string]*appliances.AirconModeTemplate{
			"auto": {
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
						Values: []interface{}{"auto", "1", "2", "3", "4"},
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
						Values: []interface{}{"auto", "1", "2", "3", "4"},
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
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
			"dry": {
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
						Values: []interface{}{"auto", "1", "2", "3", "4"},
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
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
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
						Values: []interface{}{"auto", "1", "2", "3", "4"},
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
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
			"fan": {
				Fan: &appliances.ActionTemplate{
					Type:    appliances.LIST,
					Default: "auto",
					List: &appliances.List{
						Values: []interface{}{"auto", "1", "2", "3", "4"},
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
					Default: "keep",
					List: &appliances.List{
						Values: []interface{}{"keep", "swing"},
					},
				},
			},
		},
	},
}

func (r *sc4ua) Template() *appliances.Template {
	return Template
}
