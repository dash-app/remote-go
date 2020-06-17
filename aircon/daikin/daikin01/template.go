package daikin01

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "daikin",
	Model:  "01",
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
					Type: template.RANGE,
					Default: 0.0,
					Range: &template.Range{
						Step: 0.5,
						From: -5.0,
						To: 5.0,
					},
				},
			},
			"cool": &template.AirconMode{},
			"dry":  &template.AirconMode{},
			"heat": &template.AirconMode{},
		},
	},
}

func (r *daikin01) Template() *template.Template {
	return Template
}