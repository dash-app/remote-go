package hk9817

import (
	"github.com/dash-app/remote-go/appliances"
)

type hk9817 struct {
}

func New() appliances.RemoteController {
	return &hk9817{}
}
