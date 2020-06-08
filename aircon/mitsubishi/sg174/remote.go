package sg174

import "strconv"

type Remote struct {
	Operation      bool        `json:"operation"`
	Mode           string      `json:"mode"`
	Temp           interface{} `json:"temp"`
	Fan            string      `json:"fan"`
	HorizontalVane string      `json:"horizontal_vane"`
	VerticalVane   string      `json:"vertical_vane"`
}

type State struct {
	Operation bool                  `json:"operation"`
	Mode      string                `json:"mode"`
	Entry     map[string]*ModeEntry `json:"entry"`
}

type ModeEntry struct {
	// Temp - cool/heat: float / dry: string
	Temp           interface{} `json:"temp"`
	Fan            string      `json:"fan"`
	HorizontalVane string      `json:"horizontal_vane"`
	VerticalVane   string      `json:"vertical_vane"`
}

func (s *State) ToEntry() *Remote {
	return &Remote{
		Operation:      s.Operation,
		Mode:           s.Mode,
		Temp:           s.Entry[s.Mode].Temp,
		Fan:            s.Entry[s.Mode].Fan,
		HorizontalVane: s.Entry[s.Mode].HorizontalVane,
		VerticalVane:   s.Entry[s.Mode].VerticalVane,
	}
}

func (r *Remote) Generate() ([][]int, error) {
	signal := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00},
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00},
	}

	if r.Operation {
		signal[1][5] = 0x20
	}

	if r.Mode == "cool" {
		signal[1][6] = 0x58
	} else if r.Mode == "dry" {
		signal[1][6] = 0x51
	} else if r.Mode == "heat" {
		signal[1][6] = 0x48
	} else if r.Mode == "fan" {
		signal[1][6] = 0x38
	}

	// TODO: Switch when dry
	if r.Mode == "cool" || r.Mode == "heat" {
		temp := r.Temp.(float64)
		t := int(temp) - 16
		if int(temp)*10%10 == 5 {
			t += 16
		}
		signal[1][7] = t
	}

	if r.VerticalVane == "swing" {
		signal[1][8] = 0xF8
	} else {
		v, _ := strconv.Atoi(r.VerticalVane)
		signal[1][8] = v * 16
	}

	hv := 0xC0
	if r.HorizontalVane == "swing" {
		hv += 0xF8
	} else if r.HorizontalVane != "auto" {
		h, _ := strconv.Atoi(r.HorizontalVane)
		hv += 0x08 * h
	}
	signal[1][9] = hv

	if r.Fan == "silent" {
		signal[1][9] += 0x05
	} else if r.Fan == "low" {
		signal[1][9] += 0x01
	} else if r.Fan == "mid" {
		signal[1][9] += 0x02
	} else if r.Fan == "high" {
		signal[1][9] += 0x03
	} else if r.Fan == "long" {
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
