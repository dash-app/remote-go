package rc701w

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "odelic",
	Model:  "rc701w",
	Kind:   "light",
	Light: &template.Light{
		Mode: &template.Action{
			Type:    template.LIST,
			Default: "OFF",
			List: &template.List{
				Values: []interface{}{"OFF", "ON", "NIGHT_LIGHT", "FAV"},
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
	},
}

func (r *rc701w) Template() *template.Template {
	return Template
}
