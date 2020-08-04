package aircon

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dash-app/remote-go/hex"
	"github.com/dash-app/remote-go/template"
)

type Remote interface {
	Generate(*Entry) ([]*hex.HexCode, error)
	Template() *template.Template
}

type Entry struct {
	Operation      bool        `json:"operation"`
	Mode           string      `json:"mode"`
	Temp           interface{} `json:"temp,omitempty"`
	Fan            string      `json:"fan,omitempty"`
	HorizontalVane string      `json:"horizontal_vane,omitempty"`
	VerticalVane   string      `json:"vertical_vane,omitempty"`
}

type State struct {
	Operation bool                  `json:"operation"`
	Mode      string                `json:"mode"`
	Modes     map[string]*ModeEntry `json:"modes"`
}

type ModeEntry struct {
	Temp           interface{} `json:"temp,omitempty"`
	Fan            string      `json:"fan,omitempty"`
	HorizontalVane string      `json:"horizontal_vane,omitempty"`
	VerticalVane   string      `json:"vertical_vane,omitempty"`
}

func (s *State) ToEntry() *Entry {
	return &Entry{
		Operation:      s.Operation,
		Mode:           s.Mode,
		Temp:           s.Modes[s.Mode].Temp,
		Fan:            s.Modes[s.Mode].Fan,
		HorizontalVane: s.Modes[s.Mode].HorizontalVane,
		VerticalVane:   s.Modes[s.Mode].VerticalVane,
	}
}

// UpdateFromEntry - Update State from Entry (but, values must be satisfied by template)
func (s *State) UpdateFromEntry(e *Entry, t *template.Template) (*State, error) {
	// TODO: To Replace individual Validator
	if err := e.Validate(t); err != nil {
		return nil, err
	}

	// Operation
	s.Operation = e.Operation

	// Mode
	if t.Aircon.Modes[e.Mode] == nil {
		return nil, errors.New("unexpected mode provided")
	}
	s.Mode = e.Mode

	// Temp
	if t.Aircon.Modes[e.Mode].Temp != nil {
		if temp, ok := e.Temp.(float64); ok {
			s.Modes[e.Mode].Temp = temp
		} else if temp, ok := e.Temp.(int); ok {
			s.Modes[e.Mode].Temp = temp
		} else if temp, ok := e.Temp.(string); ok {
			s.Modes[e.Mode].Temp = temp
		} else {
			return nil, errors.New("invalid temp provided")
		}
	}

	// Fan
	if t.Aircon.Modes[e.Mode].Fan != nil {
		s.Modes[e.Mode].Fan = e.Fan
	}

	// Horizontal Vane
	if t.Aircon.Modes[e.Mode].HorizontalVane != nil {
		s.Modes[e.Mode].HorizontalVane = e.HorizontalVane
	}

	// Vertical Vane
	if t.Aircon.Modes[e.Mode].VerticalVane != nil {
		s.Modes[e.Mode].VerticalVane = e.VerticalVane
	}

	return s, nil
}

func (e *Entry) Validate(t *template.Template) error {
	// Operation
	if err := t.Aircon.Operation.Validate(e.Operation); err != nil {
		return fmt.Errorf("failed validate operation: %v", err)
	}

	// Mode
	if t.Aircon.Modes[e.Mode] == nil {
		return fmt.Errorf("invalid mode provided: %v", e.Mode)
	}

	// Temp
	if t.Aircon.Modes[e.Mode].Temp != nil {
		if err := t.Aircon.Modes[e.Mode].Temp.Validate(e.Temp); err != nil {
			return fmt.Errorf("invalid temp provided: %v", err)
		}
	}

	// Fan
	if t.Aircon.Modes[e.Mode].Fan != nil {
		if err := t.Aircon.Modes[e.Mode].Fan.Validate(e.Fan); err != nil {
			return fmt.Errorf("invalid fan provided: %v", err)
		}
	}

	// Horizontal Vane
	if t.Aircon.Modes[e.Mode].HorizontalVane != nil {
		if err := t.Aircon.Modes[e.Mode].HorizontalVane.Validate(e.HorizontalVane); err != nil {
			return fmt.Errorf("invalid horizontal_vane provided: %v", err)
		}
	}

	// Vertical Vane
	if t.Aircon.Modes[e.Mode].VerticalVane != nil {
		if err := t.Aircon.Modes[e.Mode].VerticalVane.Validate(e.VerticalVane); err != nil {
			return fmt.Errorf("invalid vertical_vane provided: %v", err)
		}
	}

	return nil
}

func New(b []byte, t *template.Template) (*Entry, error) {
	ac := &Entry{}
	var base map[string]interface{}
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, err
	}

	// Operation
	operation, ok := base["operation"].(bool)
	if !ok {
		return nil, errors.New("invalid operation provided")
	}
	if err := t.Aircon.Operation.Validate(operation); err != nil {
		return nil, fmt.Errorf("invalid operation provided: %v", err)
	}
	ac.Operation = operation

	// Mode
	mode, ok := base["mode"].(string)
	if !ok {
		return nil, errors.New("invalid mode provided")
	}
	if _, ok := t.Aircon.Modes[mode]; !ok {
		return nil, fmt.Errorf("invalid mode provided: %v", mode)
	}
	ac.Mode = mode

	// Temp
	if t.Aircon.Modes[ac.Mode].Temp != nil {
		if temp, ok := base["temp"].(float64); ok {
			if err := t.Aircon.Modes[ac.Mode].Temp.Validate(temp); err != nil {
				return nil, fmt.Errorf("invalid temp provided: %v", err)
			}
			ac.Temp = temp
		} else if temp, ok := base["temp"].(string); ok {
			if err := t.Aircon.Modes[ac.Mode].Temp.Validate(temp); err != nil {
				return nil, fmt.Errorf("invalid temp provided: %v", err)
			}
			ac.Temp = temp
		} else {
			return nil, errors.New("invalid temp provided")
		}
	}

	// Fan
	if t.Aircon.Modes[ac.Mode].Fan != nil {
		fan, ok := base["fan"].(string)
		if !ok {
			return nil, errors.New("invalid fan provided")
		}
		if err := t.Aircon.Modes[ac.Mode].Fan.Validate(fan); err != nil {
			return nil, fmt.Errorf("invalid fan provided: %v", err)
		}
		ac.Fan = fan
	}

	// Horizontal Vane
	if t.Aircon.Modes[ac.Mode].HorizontalVane != nil {
		hVane, ok := base["horizontal_vane"].(string)
		if !ok {
			return nil, errors.New("invalid horizontal_vane provided")
		}
		if err := t.Aircon.Modes[ac.Mode].HorizontalVane.Validate(hVane); err != nil {
			return nil, fmt.Errorf("invalid horizontal_vane provided: %v", err)
		}
		ac.HorizontalVane = hVane
	}

	// Vertical Vane
	if t.Aircon.Modes[ac.Mode].VerticalVane != nil {
		vVane, ok := base["vertical_vane"].(string)
		if !ok {
			return nil, errors.New("invalid vertical_vane provided")
		}
		if err := t.Aircon.Modes[ac.Mode].VerticalVane.Validate(vVane); err != nil {
			return nil, fmt.Errorf("invalid vertical_vane provided: %v", err)
		}
		ac.VerticalVane = vVane
	}

	return ac, nil
}
