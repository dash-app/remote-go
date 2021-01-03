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
			Type: template.MULTIPLE,
			Multiple: []*template.Action{
				{
					Type: template.SHOT,
					Shot: &template.Shot{
						Value: "TO_DIM",
					},
				},
				{
					Type: template.SHOT,
					Shot: &template.Shot{
						Value: "TO_BRIGHT",
					},
				},
			},
		},
		Color: &template.Action{
			Type: template.MULTIPLE,
			Multiple: []*template.Action{
				{
					Type: template.SHOT,
					Shot: &template.Shot{
						Value: "TO_COOL",
					},
				},
				{
					Type: template.SHOT,
					Shot: &template.Shot{
						Value: "TO_WARM",
					},
				},
			},
		},
	},
}

func (r *ira03h) Template() *template.Template {
	return Template
}
