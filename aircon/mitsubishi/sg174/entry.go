package sg174

import (
	"strconv"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/hex"
)

func (r *sg174) Generate(e *aircon.Entry) ([]*hex.HexCode, error) {
	code := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00},
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00},
	}

	if e.Operation {
		code[1][5] = 0x20
	}

	if e.Mode == "cool" {
		code[1][6] = 0x58
	} else if e.Mode == "dry" {
		code[1][6] = 0x51
	} else if e.Mode == "heat" {
		code[1][6] = 0x48
	} else if e.Mode == "fan" {
		code[1][6] = 0x38
	}

	// TODO: Switch when dry
	if e.Mode == "cool" || e.Mode == "heat" {
		temp := e.Temp.(float32)
		t := int(temp) - 16
		if int(temp)*10%10 == 5 {
			t += 16
		}
		code[1][7] = t
	}

	if e.VerticalVane == "swing" {
		code[1][8] = 0xF8
	} else {
		v, _ := strconv.Atoi(e.VerticalVane)
		code[1][8] = v * 16
	}

	hv := 0xC0
	if e.HorizontalVane == "swing" {
		hv += 0xF8
	} else if e.HorizontalVane != "auto" {
		h, _ := strconv.Atoi(e.HorizontalVane)
		hv += 0x08 * h
	}
	code[1][9] = hv

	if e.Fan == "silent" {
		code[1][9] += 0x05
	} else if e.Fan == "low" {
		code[1][9] += 0x01
	} else if e.Fan == "mid" {
		code[1][9] += 0x02
	} else if e.Fan == "high" {
		code[1][9] += 0x03
	} else if e.Fan == "long" {
		code[1][9] += 0x03
		code[1][15] += 0x10
	}

	code[1][10] = 0x80

	// Sum
	sum := 0
	for _, c := range code[1] {
		sum += c
	}
	code[1][len(code[1])-1] = sum

	return []*hex.HexCode{
		{
			Code: code,
		},
	}, nil
}
