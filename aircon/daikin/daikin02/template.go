package daikin02

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "daikin",
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
			"auto": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 0.0,
					Range: &template.Range{
						Step: 0.5,
						From: -5.0,
						To:   5.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto"},
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
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"cool": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step: 0.5,
						From: 18.0,
						To:   32.0,
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
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"dry": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 26.0,
					Range: &template.Range{
						Step: 0.5,
						From: 18.0,
						To:   32.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto"},
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
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
			"heat": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 23.0,
					Range: &template.Range{
						Step: 0.5,
						From: 14.0,
						To:   30.0,
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
					Default: "auto",
					List: &template.List{
						Values: []interface{}{"auto", "swing", "left", "mid_left", "center", "mid_right", "right"},
					},
				},
			},
		},
	},
}

func (r *daikin02) Template() *template.Template {
	return Template
}
