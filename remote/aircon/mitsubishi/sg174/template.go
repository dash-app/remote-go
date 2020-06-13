package sg174

import "github.com/dash-app/remote-go/template"

type ModeTemplate struct {
	Temp           *template.Action `json:"temp"`
	Fan            *template.Action `json:"fan"`
	HorizontalVane *template.Action `json:"horizontal_vane"`
	VerticalVane   *template.Action `json:"vertical_vane"`
}

var Template = &template.Template{
	Vendor: "mitsubishi",
	Model:  "sg174",
	Kind:   "aircon",
	Aircon: &template.Aircon{
		Operation: &template.Action{
			Type: template.TOGGLE,
			Toggle: &template.Toggle{
				ON:  "on",
				OFF: "off",
			},
		},
		Modes: map[string]interface{}{
			"cool": &ModeTemplate{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 21.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   31.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "low", "mid", "high", "long"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "center",
					List:    []interface{}{"left", "left_mid", "center", "right_mid", "right"},
				},
			},
			"dry": &ModeTemplate{
				Temp: &template.Action{
					Type:    template.LIST,
					Default: "mid",
					List:    []interface{}{"low", "mid", "high"},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "low", "mid", "high", "long"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "center",
					List:    []interface{}{"left", "left_mid", "center", "right_mid", "right"},
				},
			},
			"heat": &ModeTemplate{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 27.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   31.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "low", "mid", "high", "long"},
				},
				HorizontalVane: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "swing", "1", "2", "3", "4", "5"},
				},
				VerticalVane: &template.Action{
					Type:    template.LIST,
					Default: "center",
					List:    []interface{}{"left", "left_mid", "center", "right_mid", "right"},
				},
			},
		},
	},
}
