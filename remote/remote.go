package remote

import "github.com/dash-app/remote-go/aircon"

type Remote interface {
	Generate(*aircon.Entry) ([]*HexCode, error)
}
type HexCode struct {
	Code     [][]int
	Interval int
}
