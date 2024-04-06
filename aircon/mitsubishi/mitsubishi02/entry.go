package mitsubishi02

import (
	"errors"
	"strconv"

	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/hex"
)

func (r *mitsubishi02) Generate(req *appliances.Request) ([]*hex.HexCode, error) {
	e := req.Aircon()
	code := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00},
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x00, 0x58, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00},
	}

	// Operation
	if e.Operation {
		code[1][5] = 0x20
	}

	// Mode
	switch e.Mode {
	case "cool":
		code[1][6] = 0x58
	case "dry":
		code[1][6] = 0x50
		code[1][7] = 0x08
	case "heat":
		code[1][6] = 0x48
	case "fan":
		code[1][6] = 0x38
	}

	// Econo
	if e.Options["econo"] != "" {
		if e.Options["econo"] == "ON" {
			code[1][6] += 0x04
		}
	}

	// Temp
	if e.Mode == "cool" || e.Mode == "heat" {
		temp := e.Temp.(float64)
		t := int(temp) - 16
		if int(temp*10)%10 == 5 {
			t += 16
		}
		code[1][7] = t
	}

	// Humid
	switch e.Mode {
	case "cool":
		code[1][8] = 0x06
	case "dry":
		switch e.Humid {
		case "40%":
			//code[1][8] = 0x00
			break
		case "50%":
			code[1][8] = 0x02
		case "60%":
			code[1][8] = 0x04
		default:
			return nil, errors.New("invalid humid provided")
		}
	}

	// Fan
	switch e.Fan {
	case "auto":
		code[1][9] = 0x00
	case "1":
		code[1][9] = 0x05
	case "2":
		code[1][9] = 0x01
	case "3":
		code[1][9] = 0x02
	case "4":
		code[1][9] = 0x03
	case "5":
		code[1][9] = 0x03
		code[1][15] = 0x10
	default:
		return nil, errors.New("invalid fan provided")
	}

	// Horizontal Vane
	switch e.HorizontalVane {
	case "auto":
		code[1][9] += 0xC0
	case "swing":
		// swing = 7 (why not 6...)
		code[1][9] += 0xC0 + (0x8 * 7)
		code[1][16] = 0x8 * 7
	case "1", "2", "3", "4", "5":
		h, _ := strconv.Atoi(e.HorizontalVane)
		code[1][9] += 0xC0 + (0x8 * h)
		code[1][16] = 0x8 * h
	default:
		return nil, errors.New("invalid horizontal_vane proided")
	}

	// Vertical Vane
	switch e.VerticalVane {
	case "swing":
		code[1][8] += 0xC0
	case "left":
		code[1][8] += 0x10
	case "mid_left":
		code[1][8] += 0x20
	case "center":
		code[1][8] += 0x30
	case "mid_right":
		code[1][8] += 0x40
	case "right":
		code[1][8] += 0x50
	default:
		return nil, errors.New("invalid vertical_vane provided")
	}

	// double beep (secret)
	// code[1][9] += 0x20 // low tone, double beep
	// code[1][9] += 0x40 // double beep
	// code[1][9] += 0x80 // low tone

	// Sum
	sum := 0
	for _, c := range code[1] {
		sum += c
	}
	code[1][len(code[1])-1] = sum & 0xFF
	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
