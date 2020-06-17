package daikin01

import "github.com/dash-app/remote-go/remote"

type daikin01 struct {
}

func New() remote.Aircon {
	return &daikin01{}
}
