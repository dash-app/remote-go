package light

import (
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
	Action string `json:"command,omitempty" example:"OFF"`
}

// State - Remote State
type State struct {
	// LastAction - Issued last action...
	LastCommand string `json:"last_action" example:"OFF"`
}
