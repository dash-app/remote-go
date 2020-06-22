package daikin03

import "github.com/dash-app/remote-go/remote"

type daikin03 struct {
}

func New() remote.Aircon {
	return &daikin03{}
}
