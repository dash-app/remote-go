package panasonic01

import "github.com/dash-app/remote-go/aircon"

type panasonic01 struct {
}

func New() aircon.Remote {
	return &panasonic01{}
}
