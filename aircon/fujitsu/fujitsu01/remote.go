package fujitsu01

import "github.com/dash-app/remote-go/appliances"

type fujitsu01 struct {
}

func New() appliances.RemoteController {
	return &fujitsu01{}
}
