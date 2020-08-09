package sg174

import "github.com/dash-app/remote-go/aircon"

type sg174 struct {
	// State Mgmt
}

func New() aircon.Remote {
	return &sg174{}
}
