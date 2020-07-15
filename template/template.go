package template

import (
	"errors"
	"fmt"
)

type Template struct {
	Vendor string  `json:"vendor"`
	Model  string  `json:"model"`
	Kind   string  `json:"kind"`
	Aircon *Aircon `json:"aircon,omitempty"`
}

type Aircon struct {
	Operation *Action                `json:"operation"`
	Modes     map[string]*AirconMode `json:"modes"`
}

type AirconMode struct {
	Temp           *Action   `json:"temp"`
	Fan            *Action   `json:"fan"`
	HorizontalVane []*Action `json:"horizontal_vane"`
	VerticalVane   []*Action `json:"vertical_vane"`
}

type Action struct {
	Type    ActionType    `json:"type"`
	Default interface{}   `json:"default,omitempty"`
	When    *When         `json:"when,omitempty"`
	List    []interface{} `json:"list"`
	Range   *Range        `json:"range"`
	Toggle  *Toggle       `json:"toggle"`
	Shot    *Shot         `json:"shot"`
}

type ActionType int32

const (
	// LIST - Multiple choice
	LIST ActionType = iota

	// RANGE - Numeric range
	RANGE

	// TOGGLE - Switch value for boolean
	TOGGLE

	// SHOT - Raise when pushed
	SHOT
)

// When - if raised (or not)
type When struct {
	Is  string `json:"is"`
	Not string `json:"not"`
}

// Range - Numeric range
type Range struct {
	Step float64 `json:"step"`
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// Toggle - Switch value for boolean
type Toggle struct {
	ON  interface{} `json:"on"`
	OFF interface{} `json:"off"`
}

// Shot - Raise when pushed
type Shot struct {
	Value interface{} `json:"value"`
}

func (a *Action) Validate(v interface{}) error {
	switch a.Type {
	case LIST:
		for _, val := range a.List {
			if val == v {
				return nil
			}
		}
		return fmt.Errorf("invalid value provided in list: %v", v)
	case RANGE:
		if val, ok := v.(float64); ok {
			if val < a.Range.From || val > a.Range.To {
				return fmt.Errorf("out of range: %v", val)
			}
			// TODO: Validate Step
		}
	case TOGGLE:
		if a.Toggle.ON != v && a.Toggle.OFF != v {
			return fmt.Errorf("invalid toggle provided: %v", v)
		}
	case SHOT:
		if a.Default != v && a.Shot.Value != v {
			return fmt.Errorf("invalid shot provided: %v", v)
		}
	default:
		return errors.New("unknown type provided")
	}

	return nil
}
