package daikin03

import "github.com/dash-app/remote-go/appliances"

type daikin03 struct {
}

func New() appliances.RemoteController {
	return &daikin03{}
}
