package panasonic01

import "github.com/dash-app/remote-go/appliances"

type panasonic01 struct {
}

func New() appliances.RemoteController {
	return &panasonic01{}
}
