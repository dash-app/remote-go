package rc701w

import "github.com/dash-app/remote-go/light"

type rc701w struct {
}

func New() light.Remote {
	return &rc701w{}
}
