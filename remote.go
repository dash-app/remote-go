package remotego

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/aircon/daikin/daikin01"
	"github.com/dash-app/remote-go/aircon/daikin/daikin02"
	"github.com/dash-app/remote-go/aircon/daikin/daikin03"
	"github.com/dash-app/remote-go/aircon/daikin/daikin04"
	"github.com/dash-app/remote-go/aircon/fujitsu/fujitsu01"
	"github.com/dash-app/remote-go/aircon/mitsubishi/mitsubishi02"
	"github.com/dash-app/remote-go/aircon/panasonic/panasonic01"
)

// TODO: Should be move to aircon package!!
type Aircon struct {
	// Remote - Generator, Template
	Remote aircon.Remote
}

type Remote struct {
	Aircon aircon.Remote
}

// AirconFromName - Get remote from vendor/model name.
func AirconFromName(vendor, model string) (*Aircon, error) {
	switch {
	case vendor == "daikin" && model == "daikin01":
		return &Aircon{
			Remote: daikin01.New(),
		}, nil
	case vendor == "daikin" && model == "daikin02":
		return &Aircon{
			Remote: daikin02.New(),
		}, nil
	case vendor == "daikin" && model == "daikin03":
		return &Aircon{
			Remote: daikin03.New(),
		}, nil
	case vendor == "daikin" && model == "daikin04":
		return &Aircon{
			Remote: daikin04.New(),
		}, nil
	case vendor == "fujitsu" && model == "fujitsu01":
		return &Aircon{
			Remote: fujitsu01.New(),
		}, nil
	case vendor == "mitsubishi" && model == "mitsubishi02":
		return &Aircon{
			Remote: mitsubishi02.New(),
		}, nil
	case vendor == "panasonic" && model == "panasonic01":
		return &Aircon{
			Remote: panasonic01.New(),
		}, nil
	}
	return nil, errors.New("no matched")
}
