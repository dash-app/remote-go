package daikin04

import "github.com/dash-app/remote-go/appliances"

type daikin04 struct {
}

func New() appliances.RemoteController {
	return &daikin04{}
}
