package fujitsu01

import "github.com/dash-app/remote-go/template"

var Template = &template.Template{
	Vendor: "fujitsu",
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
					Type:    template.RANGE,
					Default: 0.0,
					Range: &template.Range{
						Step: 0.5,
						From: 2.0,
						To:   2.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4"},
				},
				HorizontalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
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
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4"},
				},
				HorizontalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
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
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto"},
				},
				HorizontalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
					},
				},
			},
			"heat": &template.AirconMode{
				Temp: &template.Action{
					Type:    template.RANGE,
					Default: 23.0,
					Range: &template.Range{
						Step: 0.5,
						From: 16.0,
						To:   30.0,
					},
				},
				Fan: &template.Action{
					Type:    template.LIST,
					Default: "auto",
					List:    []interface{}{"auto", "1", "2", "3", "4"},
				},
				HorizontalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
					},
				},
				VerticalVane: &template.Action{
					Type:    template.MULTIPLE,
					Default: "keep",
					Multiple: []*template.Action{
						{
							Type:    template.LIST,
							Default: "keep",
							List:    []interface{}{"keep", "swing"},
						},
						{
							Type:    template.SHOT,
							Default: "keep",
							Shot: &template.Shot{
								Value: "switch",
							},
						},
					},
				},
			},
		},
	},
}

func (r *fujitsu01) Template() *template.Template {
	return Template
}
