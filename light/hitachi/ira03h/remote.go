package ira03h

import "github.com/dash-app/remote-go/light"

type ira03h struct {
}

func New() light.Remote {
	return &ira03h{}
}
