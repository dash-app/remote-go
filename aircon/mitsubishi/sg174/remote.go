package sg174

import "github.com/dash-app/remote-go/remote"

type sg174 struct {
	// State Mgmt
}

func New() remote.Aircon {
	return &sg174{}
}
