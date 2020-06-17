package remote

import (
	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/template"
)

type Aircon interface {
	Generate(*aircon.Entry) ([]*HexCode, error)
	Template() *template.Template
}

type HexCode struct {
	Code     [][]int
	Interval int
}
