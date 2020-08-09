package fujitsu01

import "github.com/dash-app/remote-go/aircon"

type fujitsu01 struct {
}

func New() aircon.Remote {
	return &fujitsu01{}
}
