package aircon

import (
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
	Operation      bool        `json:"operation" example:"false"`
	Mode           string      `json:"mode" example:"cool"`
	Temp           interface{} `json:"temp,omitempty"`
	Humid          interface{} `json:"humid,omitempty" example:"50%"`
	Fan            string      `json:"fan,omitempty" example:"auto"`
	HorizontalVane string      `json:"horizontal_vane,omitempty" example:"auto"`
	VerticalVane   string      `json:"vertical_vane,omitempty" example:"auto"`
}

type State struct {
	Operation bool                  `json:"operation" example:"false"`
	Mode      string                `json:"mode" example:"cool"`
	Modes     map[string]*ModeEntry `json:"modes"`
}

type ModeEntry struct {
	Temp           interface{} `json:"temp,omitempty"`
	Humid          interface{} `json:"humid,omitempty" example:"50%"`
	Fan            string      `json:"fan,omitempty" example:"auto"`
	HorizontalVane string      `json:"horizontal_vane,omitempty" example:"swing"`
	VerticalVane   string      `json:"vertical_vane,omitempty" example:"swing"`
}

func (s *State) ToEntry() *Entry {
	return &Entry{
		Operation:      s.Operation,
		Mode:           s.Mode,
		Temp:           s.Modes[s.Mode].Temp,
		Humid:          s.Modes[s.Mode].Humid,
		Fan:            s.Modes[s.Mode].Fan,
		HorizontalVane: s.Modes[s.Mode].HorizontalVane,
		VerticalVane:   s.Modes[s.Mode].VerticalVane,
	}
}

// DefaultState - Generate default state
func DefaultState(t *template.Template) (*State, error) {
	state := &State{}

	// Operation
	state.Operation = false

	// Mode
	if t.Aircon.Modes["cool"] != nil {
		state.Mode = "cool"
	} else {
		for k := range t.Aircon.Modes {
			state.Mode = k
			break
		}
	}

	state.Modes = make(map[string]*ModeEntry)
	for mode, modeTemplate := range t.Aircon.Modes {
		state.Modes[mode] = &ModeEntry{}
		// Temp
		if modeTemplate.Temp != nil {
			if modeTemplate.Temp.Default != nil {
				if err := modeTemplate.Temp.Validate(modeTemplate.Temp.Default); err != nil {
					return nil, err
				}
				state.Modes[mode].Temp = modeTemplate.Temp.Default
			} else {
				return nil, errors.New("temp: default value must be specified")
			}
		}

		// Humid
		if modeTemplate.Humid != nil {
			if modeTemplate.Humid.Default != nil {
				if err := modeTemplate.Humid.Validate(modeTemplate.Humid.Default); err != nil {
					return nil, err
				}
				state.Modes[mode].Humid = modeTemplate.Humid.Default
			} else {
				return nil, errors.New("humid: default value must be specified")
			}
		}

		// Fan
		if modeTemplate.Fan != nil {
			if modeTemplate.Fan.Default != nil {
				if err := modeTemplate.Fan.Validate(modeTemplate.Fan.Default); err != nil {
					return nil, err
				}
				state.Modes[mode].Fan = modeTemplate.Fan.Default.(string)
			} else {
				return nil, errors.New("fan: default value must be specified")
			}
		}

		// Horizontal Vane
		if modeTemplate.HorizontalVane != nil {
			if modeTemplate.HorizontalVane.Default != nil {
				if err := modeTemplate.HorizontalVane.Validate(modeTemplate.HorizontalVane.Default); err != nil {
					return nil, err
				}
				state.Modes[mode].HorizontalVane = modeTemplate.HorizontalVane.Default.(string)
			} else {
				return nil, errors.New("horizontal_vane: default value must be specified")
			}
		}

		// Vertical Vane
		if modeTemplate.VerticalVane != nil {
			if modeTemplate.VerticalVane.Default != nil {
				if err := modeTemplate.VerticalVane.Validate(modeTemplate.VerticalVane.Default); err != nil {
					return nil, err
				}
				state.Modes[mode].VerticalVane = modeTemplate.VerticalVane.Default.(string)
			} else {
				return nil, errors.New("vertical_vane: default value must be specified")
			}
		}
	}

	return state, nil
}

// UpdateFromEntry - Update State from Entry (but, values must be satisfied by template)
func (s *State) UpdateFromEntry(e *Entry, t *template.Template) (*State, error) {
	// TODO: To Replace individual Validator
	if err := e.Validate(t); err != nil {
		return nil, err
	}

	// Operation
	if !t.Aircon.Operation.IsShot(e.Operation) {
		s.Operation = e.Operation
	}

	// Mode
	if t.Aircon.Modes[e.Mode] == nil {
		return nil, errors.New("unexpected mode provided")
	}
	s.Mode = e.Mode

	// Temp
	if t.Aircon.Modes[e.Mode].Temp != nil && !t.Aircon.Modes[e.Mode].Temp.IsShot(e.Temp) {
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

	// Humid
	if t.Aircon.Modes[e.Mode].Humid != nil && !t.Aircon.Modes[e.Mode].Humid.IsShot(e.Humid) {
		s.Modes[e.Mode].Humid = e.Humid
	}

	// Fan
	if t.Aircon.Modes[e.Mode].Fan != nil && !t.Aircon.Modes[e.Mode].Fan.IsShot(e.Fan) {
		s.Modes[e.Mode].Fan = e.Fan
	}

	// Horizontal Vane
	if t.Aircon.Modes[e.Mode].HorizontalVane != nil && !t.Aircon.Modes[e.Mode].HorizontalVane.IsShot(e.HorizontalVane) {
		s.Modes[e.Mode].HorizontalVane = e.HorizontalVane
	}

	// Vertical Vane
	if t.Aircon.Modes[e.Mode].VerticalVane != nil && !t.Aircon.Modes[e.Mode].VerticalVane.IsShot(e.VerticalVane) {
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

	// Humid
	if t.Aircon.Modes[e.Mode].Humid != nil {
		if err := t.Aircon.Modes[e.Mode].Humid.Validate(e.Humid); err != nil {
			return fmt.Errorf("invalid humid provided: %v", err)
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
