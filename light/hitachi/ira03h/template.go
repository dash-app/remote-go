package ira03h

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "hitachi",
	Model:  "ir-a03h",
	Kind:   appliances.LIGHT,
	Light: &appliances.LightTemplate{
		Mode: &appliances.ActionTemplate{
			Type:    appliances.LIST,
			Default: "OFF",
			List: &appliances.List{
				Values: []interface{}{"OFF", "ON", "NIGHT_LIGHT", "FAV_01", "FAV_02", "FAV_03", "FAV_04"},
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
		Color: &appliances.ActionTemplate{
			Type: appliances.MULTIPLE,
			Multiple: []*appliances.ActionTemplate{
				{
					Type: appliances.SHOT,
					Shot: &appliances.Shot{
						Value: "TO_COOL",
					},
				},
				{
					Type: appliances.SHOT,
					Shot: &appliances.Shot{
						Value: "TO_WARM",
					},
				},
			},
		},
	},
}

func (r *ira03h) Template() *appliances.Template {
	return Template
}
