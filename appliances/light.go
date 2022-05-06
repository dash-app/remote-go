package appliances

import (
	"errors"
)

// Light - Call from Client
type Light struct {
	// Action - OFF, ON, NIGHT_LIGHT, ALL_LIGHT, TO_WARM, TO_COOL, TO_BRIGHT, TO_DIM
	// Direct Action (Use by when not support state-based signal)
	Action string `json:"action,omitempty" example:"OFF"`
}

// LightState - Remote State
type LightState struct {
	// LastAction - Issued last action...
	LastAction string `json:"last_action" example:"OFF"`
}

func NewLight(t *Template) (RemoteState, error) {
	return &LightState{
		LastAction: t.Light.Mode.Default.(string),
	}, nil
}

// UpdateFromState - Update(Replace) State from another state (for reset to default state)
func (s *LightState) UpdateFromState(newState *State) (RemoteState, error) {
	return newState.Action().ToState().Light, nil
}

// UpdateFromEntry - Update State from Entry (but, values must be satisfied by template)
func (s *LightState) UpdateFromEntry(req *Request, t *Template) (RemoteState, error) {
	e := req.Light()
	// Must be validate (without generate...)
	if act := t.Light.CommandToAction(e.Action); act != nil {
		if !act.IsShot(e.Action) {
			s.LastAction = e.Action
		}
	} else {
		return nil, errors.New("invalid action provided")
	}
	return s, nil
}

// ToEntry - Generate Request Entry from State
func (s *LightState) ToEntry() *Request {
	return FromLight(&Light{
		Action: s.LastAction,
	})
}

func (s *LightState) ToState() *State {
	return FromLightState(s)
}
