package daikin02

import "github.com/dash-app/remote-go/appliances"

type daikin02 struct {
}

func New() appliances.RemoteController {
	return &daikin02{}
}
