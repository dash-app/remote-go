package sg174

import "strconv"

type Entry struct {
	Operation      bool        `json:"operation"`
	Mode           string      `json:"mode"`
	Temp           interface{} `json:"temp"`
	Fan            string      `json:"fan"`
	HorizontalVane string      `json:"horizontal_vane"`
	VerticalVane   string      `json:"vertical_vane"`
}

func (e *Entry) Generate() ([][]int, error) {
	signal := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00},
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00},
	}

	if e.Operation {
		signal[1][5] = 0x20
	}

	if e.Mode == "cool" {
		signal[1][6] = 0x58
	} else if e.Mode == "dry" {
		signal[1][6] = 0x51
	} else if e.Mode == "heat" {
		signal[1][6] = 0x48
	} else if e.Mode == "fan" {
		signal[1][6] = 0x38
	}

	// TODO: Switch when dry
	if e.Mode == "cool" || e.Mode == "heat" {
		temp := e.Temp.(float64)
		t := int(temp) - 16
		if int(temp)*10%10 == 5 {
			t += 16
		}
		signal[1][7] = t
	}

	if e.VerticalVane == "swing" {
		signal[1][8] = 0xF8
	} else {
		v, _ := strconv.Atoi(e.VerticalVane)
		signal[1][8] = v * 16
	}

	hv := 0xC0
	if e.HorizontalVane == "swing" {
		hv += 0xF8
	} else if e.HorizontalVane != "auto" {
		h, _ := strconv.Atoi(e.HorizontalVane)
		hv += 0x08 * h
	}
	signal[1][9] = hv

	if e.Fan == "silent" {
		signal[1][9] += 0x05
	} else if e.Fan == "low" {
		signal[1][9] += 0x01
	} else if e.Fan == "mid" {
		signal[1][9] += 0x02
	} else if e.Fan == "high" {
		signal[1][9] += 0x03
	} else if e.Fan == "long" {
		signal[1][9] += 0x03
		signal[1][15] += 0x10
	}

	signal[1][10] = 0x80

	// Sum
	sum := 0
	for _, c := range signal[1] {
		sum += c
	}
	signal[1][len(signal[1])-1] = sum

	return signal, nil
}
