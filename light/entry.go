package light

import (
	"errors"

	"github.com/dash-app/remote-go/hex"
	"github.com/dash-app/remote-go/template"
)

type Remote interface {
	Generate(*Entry) ([]*hex.HexCode, error)
	Template() *template.Template
}

// Entry - Call from Client
type Entry struct {
	// Action - OFF, ON, NIGHT_LIGHT, ALL_LIGHT, TO_WARM, TO_COOL, TO_BRIGHT, TO_DIM
	// Direct Action (Use by when not support state-based signal)
	Action string `json:"action,omitempty" example:"OFF"`
}

// State - Remote State
type State struct {
	// LastAction - Issued last action...
	LastAction string `json:"last_action" example:"OFF"`
}

// DefaultState - Generate default state
func DefaultState(t *template.Template) (*State, error) {
	state := &State{}

	state.LastAction = t.Light.Mode.Default.(string)
	return state, nil
}

// UpdateFromEntry - Update State from Entry (but, values must be satisfied by template)
func (s *State) UpdateFromEntry(e *Entry, t *template.Template) (*State, error) {
	// Must be validate (without Generate...)
	if act := t.Light.CommandToAction(e.Action); act != nil {
		if !act.IsShot(e.Action) {
			s.LastAction = e.Action
		}
	} else {
		return nil, errors.New("invalid action provided")
	}
	return s, nil
}
