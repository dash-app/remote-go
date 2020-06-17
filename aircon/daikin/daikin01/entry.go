package daikin01

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/remote"
)

func (r *daikin01) Generate(e *aircon.Entry) ([]*remote.HexCode, error) {
	code := [][]int{
		{0x00},
		{0x11, 0xDA, 0x27, 0x00, 0xC5, 0x00, 0x00, 0xD7},
		{0x11, 0xDA, 0x27, 0x00, 0x42, 0x00, 0x00, 0x54},
		{0x11, 0xDA, 0x27, 0x00, 0x00, 0x08, 0x24, 0x00, 0xA0, 0x00, 0x00, 0x06, 0x60, 0x00, 0x00, 0xC1, 0x00, 0x00, 0x00},
	}

	// Operation
	if e.Operation {
		code[3][5] = 0x09
	}

	// Mode
	if e.Mode == "auto" {
		//code[3][5] += 0x00
		code[3][7] = 0x80
	} else if e.Mode == "cool" {
		code[3][5] += 0x30
	} else if e.Mode == "dry" {
		code[3][5] += 0x20
	} else if e.Mode == "heat" {
		code[3][5] += 0x40
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	temp := e.Temp.(float64)
	if e.Mode == "cool" || e.Mode == "heat" {
		code[3][6] = 0x24 + int((temp-18.0)*2.0)
	} else if e.Mode == "dry" || e.Mode == "auto" {
		if temp >= 0.0 {
			code[3][6] = 0xC0 + int(temp*2.0)
		} else {
			code[3][6] = 0xDF + (int(temp*2.0) + 1)
		}
	}

	// Fan
	switch e.Fan {
	case "auto":
		code[3][8] = 0xa0
	case "1":
		code[3][8] = 0x30
	case "2":
		code[3][8] = 0x40
	case "3":
		code[3][8] = 0x50
	case "4":
		code[3][8] = 0x60
	case "5":
		code[3][8] = 0x70
	default:
		return nil, errors.New("invalid fan provided")
	}

	// Horizontal Vane
	if e.HorizontalVane == "swing" {
		code[3][8] += 0x0F
	} else if e.HorizontalVane != "keep" {
		return nil, errors.New("invalid horizontal_vane provided")
	}

	// Vertical Vane
	if e.VerticalVane == "swing" {
		code[3][9] += 0x0F
	} else if e.VerticalVane != "keep" {
		return nil, errors.New("invalid vertical_vane provided")
	}

	sum := 0
	for _, c := range code[3] {
		sum += c
	}

	// choose 16 bit -> 0xNNN -> 0x_NN
	code[3][len(code[3])-1] = sum & 0x0FF

	return []*remote.HexCode{
		{
			Code: code,
		},
	}, nil
}
