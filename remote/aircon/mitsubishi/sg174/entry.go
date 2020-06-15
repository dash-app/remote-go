package sg174

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/dash-app/remote-go/remote"
)

type Entry struct {
	Operation      bool        `json:"operation"`
	Mode           string      `json:"mode"`
	Temp           interface{} `json:"temp"`
	Fan            string      `json:"fan"`
	HorizontalVane string      `json:"horizontal_vane"`
	VerticalVane   string      `json:"vertical_vane"`
}

func (e *Entry) UnmarshalJSON(b []byte) error {
	var base map[string]interface{}
	if err := json.Unmarshal(b, &base); err != nil {
		return err
	}

	// Operation
	operation, ok := base["operation"].(bool)
	if !ok {
		return errors.New("invalid operation provided")
	}
	if err := Template.Aircon.Operation.Validate(operation); err != nil {
		return fmt.Errorf("invalid operation provided: %v", err)
	}
	e.Operation = operation

	// Mode
	mode, ok := base["mode"].(string)
	if !ok {
		return errors.New("invalid mode provided")
	}
	if _, ok := Template.Aircon.Modes[mode]; !ok {
		return fmt.Errorf("invalid mode provided: %v", mode)
	}
	e.Mode = mode

	// Temp
	temp, ok := base["temp"].(float64)
	if !ok {
		return errors.New("invalid temp provided")
	}

	if err := Template.Aircon.Modes[e.Mode].(*ModeTemplate).Temp.Validate(temp); err != nil {
		return fmt.Errorf("invalid temp provided: %v", err)
	}
	e.Temp = temp

	// Fan
	fan, ok := base["fan"].(string)
	if !ok {
		return errors.New("invalid fan provided")
	}
	if err := Template.Aircon.Modes[e.Mode].(*ModeTemplate).Fan.Validate(fan); err != nil {
		return fmt.Errorf("invalid fan provided: %v", err)
	}
	e.Fan = fan

	// Horizontal Vane
	hvane, ok := base["horizontal_vane"].(string)
	if !ok {
		return errors.New("invalid horizontal_vane provided")
	}
	if err := Template.Aircon.Modes[e.Mode].(*ModeTemplate).HorizontalVane.Validate(hvane); err != nil {
		return fmt.Errorf("invalid horizontal_vane provided: %v", err)
	}
	e.HorizontalVane = hvane

	// Vertical Vane
	vvane, ok := base["vertical_vane"].(string)
	if !ok {
		return errors.New("invalid vertical_vane provided")
	}
	if err := Template.Aircon.Modes[e.Mode].(*ModeTemplate).VerticalVane.Validate(vvane); err != nil {
		return fmt.Errorf("invalid vertical_vane provided: %v", err)
	}
	e.VerticalVane = vvane

	return nil
}

func (e *Entry) Generate() ([]*remote.HexCode, error) {
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
		temp := e.Temp.(float64)
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

	return []*remote.HexCode{
		{
			Code: code,
		},
	}, nil
}
