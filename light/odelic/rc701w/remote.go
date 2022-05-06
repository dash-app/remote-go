package rc701w

import "github.com/dash-app/remote-go/appliances"

type rc701w struct {
}

func New() appliances.RemoteController {
	return &rc701w{}
}
