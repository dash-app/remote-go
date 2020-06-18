package daikin02

import "github.com/dash-app/remote-go/remote"

type daikin02 struct {
}

func New() remote.Aircon {
	return &daikin02{}
}
