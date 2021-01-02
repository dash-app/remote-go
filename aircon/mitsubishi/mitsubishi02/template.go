package mitsubishi02

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "mitsubishi",
	Model:  "02",
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
			"cool": {
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step:   0.5,
						From:   16.0,
						To:     31.0,
						Suffix: "℃",
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "1", "2", "3", "4", "5"},
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
						Values: []interface{}{"left", "mid_left", "center", "swing", "mid_right", "right"},
					},
				},
			},
			"dry": {
				Humid: &template.Action{
					Type:    template.LIST,
					Default: "50%",
					List: &template.List{
						Values: []interface{}{"40%", "50%", "60%"},
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "1", "2", "3", "4", "5"},
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
						Values: []interface{}{"left", "mid_left", "center", "swing", "mid_right", "right"},
					},
				},
			},
			"heat": {
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step:   0.5,
						From:   16.0,
						To:     31.0,
						Suffix: "℃",
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "1", "2", "3", "4", "5"},
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
						Values: []interface{}{"left", "mid_left", "center", "swing", "mid_right", "right"},
					},
				},
			},
			"fan": {
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "1", "2", "3", "4", "5"},
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
						Values: []interface{}{"left", "mid_left", "center", "swing", "mid_right", "right"},
					},
				},
			},
		},
	},
}

func (r *mitsubishi02) Template() *template.Template {
	return Template
}
