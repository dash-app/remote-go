package ira03h

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "hitachi",
	Model:  "ir-a03h",
	Kind:   "light",
	Light: &template.Light{
		Mode: &template.Action{
			Type:    template.LIST,
			Default: "OFF",
			List: &template.List{
				Values: []interface{}{"OFF", "ON", "NIGHT_LIGHT", "FAV_01", "FAV_02", "FAV_03", "FAV_04"},
			},
		},
		Brightness: &template.Action{
			Type: template.LIST,
			List: &template.List{
				Shot:   true,
				Values: []interface{}{"TO_DIM", "TO_BRIGHT"},
			},
		},
		Color: &template.Action{
			Type: template.LIST,
			List: &template.List{
				Shot:   true,
				Values: []interface{}{"TO_COOL", "TO_WARM"},
			},
		},
	},
}

func (r *ira03h) Template() *template.Template {
	return Template
}
