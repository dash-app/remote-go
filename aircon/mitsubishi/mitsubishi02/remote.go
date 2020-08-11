package mitsubishi02

import "github.com/dash-app/remote-go/aircon"

type mitsubishi02 struct {
}

func New() aircon.Remote {
	return &mitsubishi02{}
}
