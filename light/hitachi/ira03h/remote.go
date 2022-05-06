package ira03h

import (
	"github.com/dash-app/remote-go/appliances"
)

type ira03h struct {
}

func New() appliances.RemoteController {
	return &ira03h{}
}
