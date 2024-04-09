package sc4ua_test

import (
	"testing"

	"github.com/dash-app/remote-go/aircon/mitsubishi/sc4ua"
	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/test"
)

func Test_Mitsubishi02(t *testing.T) {
	rem := sc4ua.New()
	tmpl := rem.Template().Aircon

	tests := []*test.RemoteTest{
		{
			Title:  "Cool",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x69, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0x96, 0xFB, 0xFF, 0xFF},
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
			Title:  "Cool/OFF",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x34, 0x01, 0x69, 0x04, 0x00, 0x00, 0xCB, 0xFE, 0x96, 0xFB, 0xFF, 0xFF},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      false,
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x05, 0x69, 0x04, 0x00, 0x00, 0x8B, 0xFA, 0x96, 0xFB, 0xFF, 0xFF},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           tmpl.Modes["cool"].Temp.Default,
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x02, 0x69, 0x04, 0x00, 0x00, 0x8B, 0xFD, 0x96, 0xFB, 0xFF, 0xFF},
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
			Title:  "Auto",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x03, 0x69, 0x04, 0x00, 0x00, 0x8B, 0xFC, 0x96, 0xFB, 0xFF, 0xFF},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "auto",
				Temp:           tmpl.Modes["auto"].Temp.Default,
				Fan:            tmpl.Modes["auto"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["auto"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["auto"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Fan",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x00, 0x69, 0x04, 0x00, 0x00, 0x8B, 0xFF, 0x96, 0xFB, 0xFF, 0xFF},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "fan",
				Fan:            tmpl.Modes["fan"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["fan"].HorizontalVane.Default.(string),
				VerticalVane:   tmpl.Modes["fan"].VerticalVane.Default.(string),
			}),
		},
		{
			Title:  "Cool/Temp=16.0",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x60, 0x01, 0x69, 0x04, 0x00, 0x00, 0x9F, 0xFE, 0x96, 0xFB, 0xFF, 0xFF},
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x7E, 0x01, 0x69, 0x04, 0x00, 0x00, 0x81, 0xFE, 0x96, 0xFB, 0xFF, 0xFF},
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x61, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0x9E, 0xFB, 0xFF, 0xFF},
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x63, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0x9C, 0xFB, 0xFF, 0xFF},
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x65, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0x9A, 0xFB, 0xFF, 0xFF},
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x67, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0x98, 0xFB, 0xFF, 0xFF},
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
			Title:  "Cool/HorizontalVane=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x39, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0xC6, 0xFB, 0xFF, 0xFF},
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
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0x49, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0xB6, 0xFB, 0xFF, 0xFF},
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
			Title:  "Cool/VerticalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x23, 0xCB, 0x26, 0x21, 0x03},
						{0x23, 0xCB, 0x26, 0x21, 0x03, 0x74, 0x01, 0xE9, 0x04, 0x00, 0x00, 0x8B, 0xFE, 0x16, 0xFB, 0xFF, 0xFF},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           tmpl.Modes["cool"].Temp.Default,
				Fan:            tmpl.Modes["cool"].Fan.Default.(string),
				HorizontalVane: tmpl.Modes["cool"].HorizontalVane.Default.(string),
				VerticalVane:   "swing",
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
