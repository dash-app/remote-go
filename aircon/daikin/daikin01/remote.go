package daikin01

import (
	"github.com/dash-app/remote-go/appliances"
)

type daikin01 struct {
}

func New() appliances.RemoteController {
	return &daikin01{}
}
