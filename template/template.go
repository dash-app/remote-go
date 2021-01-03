package template

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Template struct {
	// Target vendor
	Vendor string `json:"vendor" example:"vendor"`

	// Target Model
	Model string `json:"model" example:"daikin01"`

	// Target Kind
	Kind string `json:"kind" example:"aircon"`

	// Aircon template
	Aircon *Aircon `json:"aircon,omitempty"`

	// Light template
	Light *Light `json:"light,omitempty"`
}

type Action struct {
	Type     ActionType  `json:"type"`
	Default  interface{} `json:"default,omitempty"`
	List     *List       `json:"list,omitempty"`
	Range    *Range      `json:"range,omitempty"`
	Toggle   *Toggle     `json:"toggle,omitempty"`
	Shot     *Shot       `json:"shot,omitempty"`
	Multiple []*Action   `json:"multiple,omitempty"`
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

	// MULTIPLE - Multiple actions
	MULTIPLE
)

// List - List entries
type List struct {
	Shot   bool          `json:"shot,omitempty"`
	Values []interface{} `json:"values"`
}

// Range - Numeric range
type Range struct {
	Step   float64 `json:"step"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
	Prefix string  `json:"prefix,omitempty"`
	Suffix string  `json:"suffix,omitempty"`
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

func (t ActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *ActionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("data should be a string, got %s", data)
	}

	switch s {
	case "LIST":
		*t = LIST
	case "RANGE":
		*t = RANGE
	case "TOGGLE":
		*t = TOGGLE
	case "SHOT":
		*t = SHOT
	case "MULTIPLE":
		*t = MULTIPLE
	default:
		return fmt.Errorf("invalid actiontype %s", s)
	}

	return nil
}

func (t ActionType) String() string {
	switch t {
	case LIST:
		return "LIST"
	case RANGE:
		return "RANGE"
	case TOGGLE:
		return "TOGGLE"
	case SHOT:
		return "SHOT"
	case MULTIPLE:
		return "MULTIPLE"
	default:
		return "UNKNOWN"
	}
}

func (a *Action) IsShot(v interface{}) bool {
	if a.Type == SHOT {
		return true
	} else if a.Type == LIST && a.List.Shot {
		return true
	} else if a.Type == MULTIPLE {
		for _, m := range a.Multiple {
			if m.IsShot(v) {
				return true
			}
		}
	}
	//if a.Type == SHOT {
	//	return true
	//} else if a.Type == MULTIPLE {
	//	for _, act := range a.Multiple {
	//		res := act.IsShot(v)
	//		if res {
	//			return true
	//		}
	//	}
	//}
	return false
}

func (a *Action) Validate(v interface{}) error {
	switch a.Type {
	case LIST:
		for _, val := range a.List.Values {
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
	case MULTIPLE:
		if v == a.Default {
			return nil
		}
		for _, m := range a.Multiple {
			if v == m.Default && m.Shot.Value != v {
				return nil
			} else if err := m.Validate(v); err == nil {
				return nil
			}
		}
		return fmt.Errorf("multiple actions couldn't satisfied: %v", v)
	default:
		return errors.New("unknown type provided")
	}

	return nil
}
