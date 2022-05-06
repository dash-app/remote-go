package mitsubishi02

import "github.com/dash-app/remote-go/appliances"

type mitsubishi02 struct {
}

func New() appliances.RemoteController {
	return &mitsubishi02{}
}
