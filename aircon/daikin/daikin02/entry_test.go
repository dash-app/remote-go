package daikin02_test

import (
	"testing"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/aircon/daikin/daikin02"
	"github.com/dash-app/remote-go/test"
)

func Test_Daikin02(t *testing.T) {
	rem := daikin02.New()

	tests := []*test.ACTestEntry{
		{
			Title:  "Cool",
			Remote: rem,
			Original: [][]int{
				{0x11, 0xDA, 0x27, 0x00, 0x02, 0xD0, 0x02, 0x00, 0x80, 0x0E, 0x46, 0x30, 0xE9, 0x1E, 0x00, 0x08, 0x00, 0x24, 0x00, 0x1D},
				{0x11, 0xDA, 0x27, 0x00, 0x00, 0x39, 0x24, 0x00, 0xA0, 0x00, 0x00, 0x06, 0x60, 0x00, 0x0A, 0xC4, 0x80, 0x24, 0xE7},
			},
			Entry: &aircon.Entry{
				Operation:      true,
				Mode:           "cool",
				Temp:           18.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			},
		},
		{
			Title:  "Heat",
			Remote: rem,
			Original: [][]int{
				{0x11, 0xDA, 0x27, 0x00, 0x02, 0xD0, 0x02, 0x00, 0x80, 0x10, 0x46, 0x30, 0xE9, 0x1E, 0x00, 0x08, 0x00, 0x24, 0x00, 0x1F},
				{0x11, 0xDA, 0x27, 0x00, 0x00, 0x49, 0x32, 0x00, 0xA0, 0x00, 0x00, 0x06, 0x60, 0x00, 0x0A, 0xC4, 0x80, 0x24, 0x05},
			},
			Entry: &aircon.Entry{
				Operation:      true,
				Mode:           "heat",
				Temp:           25.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			},
		},
		{
			Title:  "Dry",
			Remote: rem,
			Original: [][]int{
				{0x11, 0xDA, 0x27, 0x00, 0x02, 0xD0, 0x02, 0x00, 0x80, 0x0F, 0x46, 0x30, 0xE9, 0x1E, 0x00, 0x08, 0x00, 0x24, 0x00, 0x1E},
				{0x11, 0xDA, 0x27, 0x00, 0x00, 0x29, 0x32, 0x32, 0xA0, 0x00, 0x00, 0x06, 0x60, 0x00, 0x0A, 0xC4, 0x80, 0x24, 0x17},
			},
			Entry: &aircon.Entry{
				Operation:      true,
				Mode:           "dry",
				Temp:           25.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			},
		},
		{
			Title:  "OFF",
			Remote: rem,
			Original: [][]int{
				{0x11, 0xDA, 0x27, 0x00, 0x02, 0xD0, 0x82, 0x00, 0x80, 0x02, 0x46, 0xB0, 0xE9, 0x1E, 0x00, 0x08, 0x00, 0x24, 0x00, 0x11},
				{0x11, 0xDA, 0x27, 0x00, 0x00, 0x38, 0x24, 0x00, 0xA0, 0x00, 0x00, 0x06, 0x60, 0x00, 0x0A, 0xC4, 0x80, 0x24, 0xE6},
			},
			Entry: &aircon.Entry{
				Operation:      false,
				Mode:           "cool",
				Temp:           18.0,
				Fan:            "auto",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			},
		},
		{
			Title:  "Cool/Fan=5",
			Remote: rem,
			Original: [][]int{
				{0x11, 0xDA, 0x27, 0x00, 0x02, 0xD0, 0x02, 0x00, 0x80, 0x0E, 0x46, 0x30, 0xE9, 0x1E, 0x00, 0x08, 0x00, 0x24, 0x00, 0x1D},
				{0x11, 0xDA, 0x27, 0x00, 0x00, 0x39, 0x32, 0x00, 0x70, 0x00, 0x00, 0x06, 0x60, 0x00, 0x0A, 0xC4, 0x80, 0x24, 0xC5},
			},
			Entry: &aircon.Entry{
				Operation:      true,
				Mode:           "cool",
				Temp:           25.0,
				Fan:            "5",
				HorizontalVane: "auto",
				VerticalVane:   "auto",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if err := test.Validate(); err != nil {
				t.Error(err)
				t.Fail()
			}
			if err := test.Compare(); err != nil {
				t.Error(err)
				t.Fail()
			}
		})
	}
}
