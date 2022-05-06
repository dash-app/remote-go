package sg174

import "github.com/dash-app/remote-go/appliances"

type sg174 struct {
	// State Mgmt
}

func New() appliances.RemoteController {
	return &sg174{}
}
