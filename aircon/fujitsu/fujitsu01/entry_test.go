package fujitsu01_test

import (
	"testing"

	"github.com/dash-app/remote-go/aircon/fujitsu/fujitsu01"
	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/test"
)

func Test_Fujitsu01(t *testing.T) {
	rem := fujitsu01.New()

	tests := []*test.RemoteTest{
		{
			Title:  "Auto/Operation=true",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x2D},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x81, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x2C},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "auto",
				Temp:           0.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Auto/Operation=false",
			Remote: rem,
			Original: []*test.Original{
				{
					[][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0x02, 0xFD},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      false,
				Mode:           "auto",
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x1C},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x1B},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Heat/Operation=true",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x78, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x31},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x79, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x30},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           23.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Dry/Operation=true",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x18},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x17},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Auto/Operation=true/temp=2.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x12, 0x00, 0x19},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x81, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x12, 0x00, 0x18},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "auto",
				Temp:           2.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Auto/Operation=true/temp=-2.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0xEC, 0x12, 0x00, 0x41},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x81, 0x00, 0x00, 0x00, 0x00, 0x00, 0xEC, 0x12, 0x00, 0x40},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "auto",
				Temp:           -2.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/temp=30.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0xB0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xFC},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0xB1, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xFB},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           30.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/temp=18.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x50, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x5C},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x51, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x5B},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           18.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Heat/Operation=true/temp=30.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0xB0, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xF9},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0xB1, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xF8},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           30.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Heat/Operation=true/temp=16.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x40, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x69},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x41, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x68},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           16.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Dry/Operation=true/temp=30.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0xB0, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xF8},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0xB1, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xF7},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           30.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Dry/Operation=true/temp=18.0",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x50, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x58},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x51, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x57},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           18.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/Fan=1",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x1B},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x1A},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "1",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/Fan=2",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x03, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x19},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x03, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x18},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "2",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/Fan=3",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x06, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x16},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x06, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x15},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "3",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/Fan=4",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x08, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x14},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x08, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x13},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "4",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/HorizontalVane=switch",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0x6C, 0x93},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "switch",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/HorizontalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x0C},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x0B},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "swing",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Operation=true/VerticalVane=switch",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0x79, 0x86},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "switch",
			}),
		},
		{
			Title:  "Cool/Operation=true/VerticalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x20, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xFC},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x20, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xFB},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "swing",
			}),
		},
		{
			Title:  "Cool/Operation=true/HorizontalVane=swing/VerticalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					// 運転内容が変更されたとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x90, 0x01, 0x30, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xEC},
					},
				},
				{
					// オンにするとき
					Code: [][]int{
						{0x14, 0x63, 0x00, 0x10, 0x10, 0xFE, 0x0B, 0x41, 0x91, 0x01, 0x30, 0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0xEB},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "swing",
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
