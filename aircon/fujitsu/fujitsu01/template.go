package fujitsu01

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "fujitsu",
	Model:  "01",
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
						From:   2.0,
						To:     2.0,
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
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
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
						To:     30.0,
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
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
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
						To:     30.0,
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
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
					},
				},
			},
			"heat": {
				Temp: &appliances.ActionTemplate{
					Type:    appliances.RANGE,
					Default: 23.0,
					Range: &appliances.Range{
						Step:   0.5,
						From:   16.0,
						To:     30.0,
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
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &appliances.ActionTemplate{
					Type:    appliances.MULTIPLE,
					Default: "keep",
					Multiple: []*appliances.ActionTemplate{
						{
							Type:    appliances.LIST,
							Default: "keep",
							List: &appliances.List{
								Values: []interface{}{"keep", "swing"},
							},
						},
						{
							Type:    appliances.SHOT,
							Default: "keep",
							Shot: &appliances.Shot{
								Value: "switch",
							},
						},
					},
				},
			},
		},
	},
}

func (r *fujitsu01) Template() *appliances.Template {
	return Template
}
