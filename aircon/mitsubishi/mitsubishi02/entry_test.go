package mitsubishi02_test

import (
	"testing"

	"github.com/dash-app/remote-go/aircon/mitsubishi/mitsubishi02"
	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/test"
)

func Test_Mitsubishi02(t *testing.T) {
	rem := mitsubishi02.New()
	tmpl := rem.Template().Aircon

	tests := []*test.RemoteTest{
		{
			Title:  "Cool",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x9D},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Dry",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x50, 0x08, 0x32, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x8F},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Humid:          tmpl.Modes["dry"].Humid.Default.(string),
				Fan:            tmpl.Modes["dry"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["dry"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["dry"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Heat",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x48, 0x0A, 0x30, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x87},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           tmpl.Modes["heat"].Temp.Default,
				Fan:            tmpl.Modes["heat"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["heat"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["heat"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Temp=16.0",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x00, 0x36, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x93},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Range.From,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Temp=31.0",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0F, 0x36, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0xA2},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Range.To,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Fan=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC5, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0xA2},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            "1",
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Fan=2",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC1, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x9E},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            "2",
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Fan=3",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC2, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x9F},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            "3",
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Fan=4",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC3, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0xA0},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            "4",
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Fan=5",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC3, 0x00, 0x00, 0x00, 0x00, 0x10, 0x10, 0x00, 0xB0},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            "5",
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xC8, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x08, 0xAD},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: "1",
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/HorizontalVane=Swing",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x38, 0x0D},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: "swing",
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Fan=1,HorizontalVane=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x36, 0xCD, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x08, 0xB2},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            "1",
				HorizontalVane: "1",
				VerticalVane:   tmpl.Modes["cool"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/VerticalVane=left",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x01, 0x00},
						{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x58, 0x0A, 0x16, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x7D},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   "left",
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if err := test.Compare(); err != nil {
				t.Error(err)
				t.Fail()
			}
		})
	}
}
