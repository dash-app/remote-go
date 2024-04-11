package hk9817

import (
	"github.com/dash-app/remote-go/appliances"
)

var Template = &appliances.Template{
	Vendor: "panasonic",
	Model:  "hk9817",
	Kind:   appliances.LIGHT,
	Light: &appliances.LightTemplate{
		Mode: &appliances.ActionTemplate{
			Type:    appliances.LIST,
			Default: "OFF",
			List: &appliances.List{
				Values: []interface{}{"OFF", "ON", "NIGHT_LIGHT", "FAV", "COOL", "WARM", "DIM"},
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

func (r *hk9817) Template() *appliances.Template {
	return Template
}
