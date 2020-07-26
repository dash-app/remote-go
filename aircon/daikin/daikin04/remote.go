package daikin04

import "github.com/dash-app/remote-go/aircon"

type daikin04 struct {
}

func New() aircon.Remote {
	return &daikin04{}
}
