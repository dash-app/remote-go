package daikin03

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/remote"
)

func (r *daikin03) Generate(e *aircon.Entry) ([]*remote.HexCode, error) {
	code := [][]int{
		{0x00},
		{0x11, 0xDA, 0x27, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x11, 0xDA, 0x27, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x60, 0x00, 0x00, 0xC3, 0x00, 0x00, 0x00},
	}

	// Operation
	if e.Operation {
		code[2][5] = 0x09
	} else {
		code[2][5] = 0x08
		code[1][11] = 0x80
	}

	// Mode
	if e.Mode == "auto" {
		//code[2][5] += 0x00
		code[2][7] = 0x80
	} else if e.Mode == "cool" {
		code[2][5] += 0x30
	} else if e.Mode == "dry" {
		code[2][5] += 0x20
		code[2][7] = 0x80
	} else if e.Mode == "heat" {
		code[2][5] += 0x40
	} else {
		return nil, errors.New("invalid mode provided")
	}

	// Temp
	temp := e.Temp.(float64)
	if e.Mode == "cool" || e.Mode == "heat" {
		code[2][6] = 0x24 + int((temp-18.0)*2.0)
	} else if e.Mode == "dry" || e.Mode == "auto" {
		if temp >= 0.0 {
			code[2][6] = 0xC0 + int(temp*2.0)
		} else {
			code[2][6] = 0xDF + (int(temp*2.0) + 1)
		}
	}

	// Fan
	switch e.Fan {
	case "auto":
		code[2][8] = 0xA0
	case "1":
		code[2][8] = 0x30
	case "2":
		code[2][8] = 0x40
	case "3":
		code[2][8] = 0x50
	case "4":
		code[2][8] = 0x60
	case "5":
		code[2][8] = 0x70
	default:
		return nil, errors.New("invalid fan provided")
	}

	// Horizontal Vane
	switch e.HorizontalVane {
	case "auto":
		code[1][12] = 0xE0
	case "swing":
		code[1][12] = 0x00
		code[2][8] += 0x0F
	case "1":
		code[1][12] = 0x10
	case "2":
		code[1][12] = 0x20
	case "3":
		code[1][12] = 0x30
	case "4":
		code[1][12] = 0x40
	case "5":
		code[1][12] = 0x50
	}

	if e.VerticalVane == "swing" {
		code[2][9] += 0x0F
	}

	sum1 := 0
	for _, c := range code[1] {
		sum1 += c
	}

	sum2 := 0
	for _, c := range code[2] {
		sum2 += c
	}

	// choose 16 bit -> 0xNNN -> 0x_NN
	code[1][len(code[1])-1] = sum1 & 0x0FF
	code[2][len(code[2])-1] = sum2 & 0x0FF

	return []*remote.HexCode{
		{
			Code: code,
		},
	}, nil
}
