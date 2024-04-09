package sc4ua

import (
	"errors"

	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/hex"
)

func (r *sc4ua) Generate(req *appliances.Request) ([]*hex.HexCode, error) {
	e := req.Aircon()
	code := [][]int{
		{0x23, 0xCB, 0x26, 0x21, 0x03},
		{0x23, 0xCB, 0x26, 0x21, 0x03, 0x60, 0x01, 0x00, 0x04, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF},
	}

	// Operation
	if e.Operation {
		code[1][5] = 0x60
		code[1][11] -= 0x60
	} else {
		code[1][5] = 0x20
		code[1][11] -= 0x20
	}

	// Mode
	switch e.Mode {
	case "auto":
		code[1][6] = 0x03
		code[1][12] = 0xFC
	case "cool":
		code[1][6] = 0x01
		code[1][12] = 0xFE
	case "dry":
		code[1][6] = 0x05
		code[1][12] = 0xFA
	case "heat":
		code[1][6] = 0x02
		code[1][12] = 0xFD
	case "fan":
		code[1][6] = 0x00
		code[1][12] = 0xFF
	}

	// Temp
	if e.Mode == "auto" || e.Mode == "cool" || e.Mode == "dry" || e.Mode == "heat" {
		temp := e.Temp.(float64)
		t := (int(temp*10) - 160) / 5
		code[1][5] += t
		code[1][11] -= t
	} else {
		// code[1][4] = 0x00
		code[1][5] = 0x74
		code[1][11] = 0xFF - 0x74
	}

	// Fan
	switch e.Fan {
	case "auto":
		code[1][7] += 0x09
		code[1][13] -= 0x09
	case "1":
		code[1][7] += 0x01
		code[1][13] -= 0x01
	case "2":
		code[1][7] += 0x03
		code[1][13] -= 0x03
	case "3":
		code[1][7] += 0x05
		code[1][13] -= 0x05
	case "4":
		code[1][7] += 0x07
		code[1][13] -= 0x07
	default:
		return nil, errors.New("invalid fan provided")
	}

	// Horizontal Vane
	switch e.HorizontalVane {
	case "auto":
		code[1][7] += 0x60
		code[1][13] -= 0x60
	case "swing":
		code[1][7] += 0x40
		code[1][13] -= 0x40
	case "1":
		code[1][7] += 0x30
		code[1][13] -= 0x30
	case "2":
		code[1][7] += 0x50
		code[1][13] -= 0x50
	case "3":
		code[1][7] += 0x20
		code[1][13] -= 0x20
	case "4":
		code[1][7] += 0x10
		code[1][13] -= 0x10
	case "5":
		code[1][7] += 0x00
		code[1][13] -= 0x00

	default:
		return nil, errors.New("invalid horizontal_vane proided")
	}

	// Vertical Vane
	switch e.VerticalVane {
	case "keep":
		// code[1][7] += 0x00
		// code[1][13] -= 0x00
	case "swing":
		code[1][7] += 0x80
		code[1][13] -= 0x80
	default:
		return nil, errors.New("invalid vertical_vane provided")
	}

	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
