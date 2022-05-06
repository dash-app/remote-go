package daikin04_test

import (
	"testing"

	"github.com/dash-app/remote-go/aircon/daikin/daikin04"
	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/test"
)

func Test_Daikin04(t *testing.T) {
	rem := daikin04.New()

	tests := []*test.RemoteTest{
		{
			Title:  "Auto",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x11, 0x00, 0x00, 0x00, 0x80, 0x0A, 0x08, 0x88},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "auto",
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x48},
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
			Title:  "Cool/Operation=off",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x30, 0x00, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x47},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      false,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Temp=18",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x10, 0x0A, 0x08, 0x38},
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
			Title:  "Cool/Temp=32",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x2C, 0x0A, 0x08, 0x54},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           32.0,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Fan=auto",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x48},
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
			Title:  "Cool/Fan=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x03, 0x08, 0x41},
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
			Title:  "Cool/Fan=2",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x04, 0x08, 0x42},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "2",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Fan=3",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x05, 0x08, 0x43},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "3",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Fan=4",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x06, 0x08, 0x44},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "4",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/Fan=5",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0x00, 0x00, 0x00, 0x20, 0x07, 0x08, 0x45},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "5",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0xF1, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x39},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "auto",
				HorizontalVane: "swing",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Cool/VerticalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0xF2, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x3A},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "swing",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=swing,VerticalVane=swing",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x31, 0xF7, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x3F},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26,
				Fan:            "auto",
				HorizontalVane: "swing",
				VerticalVane:   "swing",
			}),
		},
		{
			Title:  "Dry",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x21, 0x00, 0x00, 0x00, 0x80, 0x0A, 0x08, 0x98},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
			}),
		},
		{
			Title:  "Heat",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
						{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x41, 0x00, 0x00, 0x00, 0x1A, 0x0A, 0x08, 0x52},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           23,
				Fan:            "auto",
				HorizontalVane: "keep",
				VerticalVane:   "keep",
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
