package daikin01

import "github.com/dash-app/remote-go/aircon"

type daikin01 struct {
}

func New() aircon.Remote {
	return &daikin01{}
}
