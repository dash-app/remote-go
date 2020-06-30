package panasonic01

import "github.com/dash-app/remote-go/remote"

type panasonic01 struct {
}

func New() remote.Aircon {
	return &panasonic01{}
}
