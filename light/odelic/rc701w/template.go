package rc701w

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "odelic",
	Model:  "rc701w",
	Kind:   appliances.LIGHT,
	Light: &appliances.LightTemplate{
		Mode: &appliances.ActionTemplate{
			Type:    appliances.LIST,
			Default: "OFF",
			List: &appliances.List{
				Values: []interface{}{"OFF", "ON", "NIGHT_LIGHT", "FAV"},
			},
		},
		Brightness: &appliances.ActionTemplate{
			Type: appliances.MULTIPLE,
			Multiple: []*appliances.ActionTemplate{
				{
					Type: appliances.SHOT,
					Shot: &appliances.Shot{
						Value: "TO_DIM",
					},
				},
				{
					Type: appliances.SHOT,
					Shot: &appliances.Shot{
						Value: "TO_BRIGHT",
					},
				},
			},
		},
	},
}

func (r *rc701w) Template() *appliances.Template {
	return Template
}
