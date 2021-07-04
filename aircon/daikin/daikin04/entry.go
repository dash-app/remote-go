package daikin04

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/hex"
)

func (r *daikin04) Generate(e *aircon.Entry) ([]*hex.HexCode, error) {
	code := [][]int{
		{0x11, 0xDA, 0x27, 0xF0, 0x0D, 0x00, 0x0F},
		{0x11, 0xDA, 0x27, 0x00, 0xD3, 0x00, 0x00, 0x00, 0x00, 0x20, 0x0A, 0x08, 0x00},
	}

	// Operation
	if e.Operation {
		code[1][5] = 0x01
	} else {
		code[1][5] = 0x00
	}

	// Mode
	if e.Mode == "auto" {
		code[1][5] += 0x10
		code[1][9] = 0x80
	} else if e.Mode == "cool" {
		code[1][5] += 0x30
	} else if e.Mode == "dry" {
		code[1][5] += 0x20
		code[1][9] = 0x80
	} else if e.Mode == "heat" {
		code[1][5] += 0x40
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	if e.Mode == "cool" || e.Mode == "heat" {
		var temp int
		if t, ok := e.Temp.(float64); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp provided")
		}

		code[1][9] = 0x10 + int((temp-18)*2.0)
	}

	// Fan
	switch e.Fan {
	case "auto":
		code[1][10] = 0x0A
	case "1":
		code[1][10] = 0x03
	case "2":
		code[1][10] = 0x04
	case "3":
		code[1][10] = 0x05
	case "4":
		code[1][10] = 0x06
	case "5":
		code[1][10] = 0x07
	default:
		return nil, errors.New("invalid fan provided")
	}

	if e.HorizontalVane == "swing" && e.VerticalVane != "swing" {
		code[1][6] += 0xF1
	} else if e.HorizontalVane != "swing" && e.VerticalVane == "swing" {
		code[1][6] += 0xF2
	} else if e.HorizontalVane == "swing" && e.VerticalVane == "swing" {
		code[1][6] += 0xF7
	} else if e.HorizontalVane != "keep" && e.VerticalVane != "keep" {
		return nil, errors.New("invalid horizontal_vane / vertical_vane provided")
	}

	sum := 0
	for _, c := range code[1] {
		sum += c
	}

	// choose 16 bit -> 0xNNN -> 0x_NN
	code[1][len(code[1])-1] = sum & 0x0FF

	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
