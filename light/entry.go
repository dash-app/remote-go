package light

import (
	"text/template"

	"github.com/dash-app/remote-go/hex"
)

type Remote interface {
	Generate(*Entry) ([]*hex.HexCode, error)
	Template() *template.Template
}

// Entry - Call from Client
type Entry struct {
	// Command - OFF, ON, NIGHT_LIGHT, ALL_LIGHT, TO_WARM, TO_COOL, TO_BRIGHT, TO_DIM
	Command string `json:"command" example:"OFF"`
}

// State - Remote State
type State struct {
	// LastCommand - Issued last commands...
	LastCommand string `json:"last_command" example:"OFF"`
}
