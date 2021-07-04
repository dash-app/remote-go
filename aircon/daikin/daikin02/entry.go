package daikin02

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/hex"
)

func (r *daikin02) Generate(e *aircon.Entry) ([]*hex.HexCode, error) {
	code := [][]int{
		{0x11, 0xDA, 0x27, 0x00, 0x02, 0xD0, 0x02, 0x00, 0x80, 0x0E, 0x46, 0x30, 0xE9, 0x1E, 0x00, 0x08, 0x00, 0x24, 0x00, 0x00},
		{0x11, 0xDA, 0x27, 0x00, 0x00, 0x00, 0x24, 0x00, 0xA0, 0x00, 0x00, 0x06, 0x60, 0x00, 0x0A, 0xC4, 0x80, 0x24, 0x00},
	}

	// Operation
	if e.Operation {
		code[1][5] = 0x09
	} else {
		code[0][6] = 0x82
		code[0][9] = 0x02
		code[0][11] = 0xB0
		code[1][5] = 0x08
	}

	// Mode
	if e.Mode == "auto" {
		if e.Operation {
			code[0][9] = 0x0D
		}
		//code[1][5] += 0x00
		code[1][7] = 0x80
	} else if e.Mode == "cool" {
		if e.Operation {
			code[0][9] = 0x0E
		}
		code[1][5] += 0x30
	} else if e.Mode == "dry" {
		if e.Operation {
			code[0][9] = 0x0F
		}
		code[1][5] += 0x20
	} else if e.Mode == "heat" {
		if e.Operation {
			code[0][9] = 0x10
		}
		code[1][5] += 0x40
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	temp := e.Temp.(float64)
	if e.Mode == "cool" || e.Mode == "heat" || e.Mode == "dry" {
		code[1][6] = 0x24 + int((temp-18.0)*2.0)
		if e.Mode == "dry" {
			// Humid (60: 0x32, 50: 0x3C)
			code[1][7] = 0x32
		}
	} else if e.Mode == "auto" {
		if temp >= 0.0 {
			code[1][6] = 0xC0 + int(temp*2.0)
		} else {
			code[1][6] = 0xDF + (int(temp*2.0) + 1)
		}
	}

	// Fan
	switch e.Fan {
	case "auto":
		code[1][8] = 0xA0
	case "1":
		code[1][8] = 0x30
	case "2":
		code[1][8] = 0x40
	case "3":
		code[1][8] = 0x50
	case "4":
		code[1][8] = 0x60
	case "5":
		code[1][8] = 0x70
	default:
		return nil, errors.New("invalid fan provided")
	}

	// Horizontal Vane
	switch e.HorizontalVane {
	case "auto":
		code[0][12] = 0xE9
	case "swing":
		code[0][12] = 0xF9
	case "1":
		code[0][12] = 0x19
	case "2":
		code[0][12] = 0x29
	case "3":
		code[0][12] = 0x39
	case "4":
		code[0][12] = 0x49
	case "5":
		code[0][12] = 0x59
	default:
		return nil, errors.New("invalid horizontal_vane provided")
	}

	// Vertical Vane
	switch e.VerticalVane {
	case "auto":
		code[0][13] = 0x1E
	case "swing":
		code[0][13] = 0x1F
	case "left":
		code[0][13] = 0x09
	case "mid_left":
		code[0][13] = 0x0A
	case "center":
		code[0][13] = 0x0B
	case "mid_right":
		code[0][13] = 0x0C
	case "right":
		code[0][13] = 0x0D
	default:
		return nil, errors.New("invalid vertical_vane provided")
	}

	sum0 := 0
	sum1 := 0
	for _, c := range code[0] {
		sum0 += c
	}

	for _, c := range code[1] {
		sum1 += c
	}

	// choose 16 bit -> 0xNNN -> 0x_NN
	code[0][len(code[0])-1] = sum0 & 0x0FF
	code[1][len(code[1])-1] = sum1 & 0x0FF

	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
