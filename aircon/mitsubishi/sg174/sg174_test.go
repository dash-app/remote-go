package sg174_test

import (
	"testing"

	"github.com/dash-home/remote-go/aircon/mitsubishi/sg174"
	"github.com/dash-home/remote-go/hex"
)

func TestExampleSuccess(t *testing.T) {
	entry := sg174.Remote{
		Operation:      true,
		Mode:           "heat",
		Temp:           24.5,
		Fan:            "auto",
		HorizontalVane: "auto",
		VerticalVane:   "3",
	}
	r, err := entry.Generate()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	test_case := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00},
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x20, 0x48, 0x18, 0x30, 0xC0, 0x80, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x315},
	}
	hex.Print(test_case)
	hex.Print(r)
}
