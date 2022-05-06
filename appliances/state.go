package appliances

import (
	"errors"
)

// State - state of remote-go controller
type State struct {
	Aircon *AirconState `json:"aircon,omitempty"`
	Light  *LightState  `json:"light,omitempty"`
}

type RemoteState interface {
	UpdateFromState(*State) (RemoteState, error)
	UpdateFromEntry(*Request, *Template) (RemoteState, error)
	ToEntry() *Request
	ToState() *State
}

// NewState - Create new state from default entry
func NewState(t *Template) (RemoteState, error) {
	if t.Kind == AIRCON {
		return NewAircon(t)
	}
	if t.Kind == LIGHT {
		return NewLight(t)
	}

	return nil, errors.New("Unsupported kind")
}

func GetState(s *State) RemoteState {
	if s.Aircon != nil {
		return s.Aircon
	}
	if s.Light != nil {
		return s.Light
	}
	return nil
}

func (s *State) Action() RemoteState {
	if s.Aircon != nil {
		return s.Aircon
	}
	if s.Light != nil {
		return s.Light
	}
	return nil
}

// FromAirconState - Combine state from AC
func FromAirconState(s *AirconState) *State {
	return &State{Aircon: s}
}

// FromLightState - Combine state from Light
func FromLightState(s *LightState) *State {
	return &State{Light: s}
}
