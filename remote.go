package remotego

import (
	"errors"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/aircon/daikin/daikin01"
	"github.com/dash-app/remote-go/aircon/daikin/daikin02"
	"github.com/dash-app/remote-go/aircon/daikin/daikin03"
	"github.com/dash-app/remote-go/aircon/daikin/daikin04"
	"github.com/dash-app/remote-go/aircon/panasonic/panasonic01"
)

type Aircon struct {
	// Remote
	// TODO: Merge remote & template
	// Template
}

// AirconFromName - Get remote from vendor/model name.
func AirconFromName(vendor, model string) (aircon.Remote, error) {
	switch {
	case vendor == "daikin" && model == "daikin01":
		return daikin01.New(), nil
	case vendor == "daikin" && model == "daikin02":
		return daikin02.New(), nil
	case vendor == "daikin" && model == "daikin03":
		return daikin03.New(), nil
	case vendor == "daikin" && model == "daikin04":
		return daikin04.New(), nil
	case vendor == "panasonic" && model == "panasonic01":
		return panasonic01.New(), nil
	}
	return nil, errors.New("no matched")
}
