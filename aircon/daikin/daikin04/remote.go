package daikin04

import "github.com/dash-app/remote-go/remote"

type daikin04 struct {
}

func New() remote.Aircon {
	return &daikin04{}
}
