package panasonic01

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/hex"
)

func (r *panasonic01) Generate(e *aircon.Entry) ([]*hex.HexCode, error) {

	code := [][]int{
		{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x06},
		{0x02, 0x20, 0xE0, 0x04, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x06, 0x60, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
	}

	// Operation
	if e.Operation {
		code[1][5] = 0x01
	}

	// Mode
	if e.Mode == "auto" {
		code[1][6] = 0x34
	} else if e.Mode == "cool" {
		code[1][5] += 0x30
		code[1][6] = 0x34
	} else if e.Mode == "dry" {
		code[1][5] += 0x20
		code[1][6] = 0x34
	} else if e.Mode == "heat" {
		code[1][5] += 0x40
		code[1][6] = 0x2E
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	if e.Mode == "auto" || e.Mode == "cool" || e.Mode == "dry" || e.Mode == "heat" {
		var temp int
		if t, ok := e.Temp.(float32); ok {
			temp = int(t)
		} else if t, ok := e.Temp.(int); ok {
			temp = t
		} else {
			return nil, errors.New("invalid temp provided")
		}
		code[1][6] = 0x20 + int((temp-16)*2.0)
	} else {
		return nil, errors.New("invalid temp provided")
	}

	//Fan
	switch e.Fan {
	case "auto":
		code[1][8] = 0xAF
		code[1][13] = 0x40
	case "1":
		code[1][8] = 0x3F
		code[1][13] = 0x60
	case "2":
		code[1][8] = 0x3F
		code[1][13] = 0x40
	case "3":
		code[1][8] = 0x4F
		code[1][13] = 0x40
	case "4":
		code[1][8] = 0x5F
		code[1][13] = 0x40
	case "5":
		code[1][8] = 0x7F
		code[1][13] = 0x40
	default:
		return nil, errors.New("invalid fan provaided")
	}

	//HorizontalVane
	switch e.HorizontalVane {
	case "auto":
		code[1][8] = (code[1][8] & 0xF0) + 0x0F
	case "1":
		code[1][8] = (code[1][8] & 0xF0) + 0x01
	case "2":
		code[1][8] = (code[1][8] & 0xF0) + 0x02
	case "3":
		code[1][8] = (code[1][8] & 0xF0) + 0x03
	case "4":
		code[1][8] = (code[1][8] & 0xF0) + 0x04
	case "5":
		code[1][8] = (code[1][8] & 0xF0) + 0x05
	default:
		return nil, errors.New("invalid horizontal_vane provided")
	}

	//VerticalVane
	switch e.VerticalVane {
	case "auto":
		code[1][9] = 0x0D
	case "left":
		code[1][9] = 0x09
	case "mid_left":
		code[1][9] = 0x0A
	case "center":
		code[1][9] = 0x06
	case "mid_right":
		code[1][9] = 0x0B
	case "right":
		code[1][9] = 0x0C
	default:
		return nil, errors.New("invalid vertical_vane provided")
	}

	sum := 0
	for _, c := range code[1] {
		sum += c
	}

	// Create of check_sum
	// choose 8 bit -> 0xNNN -> 0x_NN
	code[1][len(code[1])-1] = sum & 0x0FF

	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
