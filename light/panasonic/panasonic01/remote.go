package panasonic01

import "github.com/dash-app/remote-go/light"

type panasonic01 struct {
}

func New() light.Remote {
	return &panasonic01{}
}
