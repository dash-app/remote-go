package daikin03

import "github.com/dash-app/remote-go/aircon"

type daikin03 struct {
}

func New() aircon.Remote {
	return &daikin03{}
}
