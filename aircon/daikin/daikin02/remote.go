package daikin02

import "github.com/dash-app/remote-go/aircon"

type daikin02 struct {
}

func New() aircon.Remote {
	return &daikin02{}
}
