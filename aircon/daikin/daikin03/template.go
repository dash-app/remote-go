package daikin03

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "daikin",
	Model:  "03",
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
					List:    []interface{}{"auto"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"keep", "swing"},
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
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "swing",
					List:    []interface{}{"swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "keep",
					List:    []interface{}{"keep", "swing"},
				},
			},
			"dry": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 0.0,
					Range: &template.Range{
						Step: 0.5,
						From: -3.0,
						To:   0.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "keep",
					List:    []interface{}{"keep", "swing"},
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
					List:    []interface{}{"auto", "1", "2", "3", "4", "5"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "keep",
					List:    []interface{}{"keep", "swing"},
				},
			},
		},
	},
}

func (r *daikin03) Template() *template.Template {
	return Template
}
