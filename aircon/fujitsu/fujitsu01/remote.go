package fujitsu01

import "github.com/dash-app/remote-go/remote"

type fujitsu01 struct {
}

func New() remote.Aircon {
	return &fujitsu01{}
}
