package sc4ua

import "github.com/dash-app/remote-go/appliances"

type sc4ua struct {
}

func New() appliances.RemoteController {
	return &sc4ua{}
}
