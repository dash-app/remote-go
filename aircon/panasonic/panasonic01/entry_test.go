package panasonic01_test

import (
	"testing"

	"github.com/dash-app/remote-go/aircon/panasonic/panasonic01"
	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/test"
)

func Test_Panasonic01(t *testing.T) {
	rem := panasonic01.New()

	tests := []*test.RemoteTest{
		{
			Title:  "Auto",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x01, 0x34, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x1E},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "auto",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Operation=true",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4E},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Operation=false",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x30, 0x34, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4D},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      false,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Temp=16",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x20, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x3A},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           16.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Temp=30",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x3C, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x56},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           30.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Fan=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x3F, 0x0D, 0x00, 0x06, 0x60, 0x60, 0x01, 0x00, 0x00, 0x00, 0xFE},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "1",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Fan=2",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x3F, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0xDE},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "2",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Fan=3",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x4F, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0xEE},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "3",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Fan=4",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x5F, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0xFE},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "4",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/Fan=5",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x7F, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x1E},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "5",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xA1, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x40},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "1",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=2",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xA2, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x41},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "2",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=3",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xA3, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x42},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "3",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=4",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xA4, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x43},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "4",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=5",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xA5, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x44},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "5",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1,Fan=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x31, 0x0D, 0x00, 0x06, 0x60, 0x60, 0x01, 0x00, 0x00, 0x00, 0xF0},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "1",
				HorizontalVane: "1",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1,Fan=2",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x31, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0xD0},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "2",
				HorizontalVane: "1",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1,Fan=3",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x41, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0xE0},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "3",
				HorizontalVane: "1",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1,Fan=4",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x51, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0xF0},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "4",
				HorizontalVane: "1",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/HorizontalVane=1,Fan=5",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0x71, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x10},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "5",
				HorizontalVane: "1",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Cool/VerticalVane=1",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xAF, 0x09, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4A},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "left",
			}),
		},
		{
			Title:  "Cool/VerticalVane=2",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xAF, 0x0A, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4B},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "mid_left",
			}),
		},
		{
			Title:  "Cool/VerticalVane=3",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xAF, 0x06, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x47},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "center",
			}),
		},
		{
			Title:  "Cool/VerticalVane=4",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xAF, 0x0B, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4C},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "mid_right",
			}),
		},
		{
			Title:  "Cool/VerticalVane=5",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x31, 0x34, 0x80, 0xAF, 0x0C, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4D},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "cool",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "right",
			}),
		},
		{
			Title:  "Dry",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x21, 0x34, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x3E},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           26.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Heat",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x41, 0x2E, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x58},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           23.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Heat/Temp=16",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x41, 0x20, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x4A},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           16.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Heat/Temp=30",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x41, 0x3C, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x66},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "heat",
				Temp:           30.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Dry/Temp=16",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x21, 0x20, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x2A},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           16.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			}),
		},
		{
			Title:  "Dry/Temp=30",
			Remote: rem,
			Original: []*test.Original{
				{
					Code: [][]int{
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
						{0x02, 0x20, 0xE0, 0x04, 0x00, 0x21, 0x3C, 0x80, 0xAF, 0x0D, 0x00, 0x06, 0x60, 0x40, 0x01, 0x00, 0x00, 0x00, 0x46},
					},
				},
			},
			Request: appliances.FromAircon(&appliances.Aircon{
				Operation:      true,
				Mode:           "dry",
				Temp:           30.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
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
