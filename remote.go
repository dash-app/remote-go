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

// TODO: Should be move to aircon package!!
type Aircon struct {
	// Remote - Generator, Template
	Remote aircon.Remote
}

func (ac *Aircon) DefaultState() (*aircon.State, error) {
	t := ac.Remote.Template()
	state := &aircon.State{}

	// Operation
	state.Operation = false

	// Mode
	if t.Aircon.Modes["cool"] != nil {
		state.Mode = "cool"
	} else {
		for k := range t.Aircon.Modes {
			state.Mode = k
			break
		}
	}

	state.Modes = make(map[string]*aircon.ModeEntry)
	for mode, modeTemplate := range t.Aircon.Modes {
		state.Modes[mode] = &aircon.ModeEntry{}
		// Temp
		if modeTemplate.Temp != nil {
			if temp, ok := modeTemplate.Temp.Default.(float64); ok {
				state.Modes[mode].Temp = temp
			} else if temp, ok := modeTemplate.Temp.Default.(int); ok {
				state.Modes[mode].Temp = temp
			} else if temp, ok := modeTemplate.Temp.Default.(string); ok {
				state.Modes[mode].Temp = temp
			} else {
				return nil, errors.New("invalid temp provided")
			}
		}

		// Fan
		if modeTemplate.Fan != nil {
			if fan, ok := modeTemplate.Fan.Default.(string); ok {
				state.Modes[mode].Fan = fan
			} else {
				return nil, errors.New("invalid fan provided")
			}
		}

		// Horizontal Vane
		if modeTemplate.HorizontalVane != nil {
			if hVane, ok := modeTemplate.HorizontalVane.Default.(string); ok {
				state.Modes[mode].HorizontalVane = hVane
			} else {
				return nil, errors.New("invalid horizontal_vane provided")
			}
		}

		// Vertical Vane
		if modeTemplate.VerticalVane != nil {
			if vVane, ok := modeTemplate.VerticalVane.Default.(string); ok {
				state.Modes[mode].VerticalVane = vVane
			} else {
				return nil, errors.New("invalid vertical_vane provided")
			}
		}
	}

	return state, nil
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
	case vendor == "panasonic" && model == "panasonic01":
		return &Aircon{
			Remote: panasonic01.New(),
		}, nil
	}
	return nil, errors.New("no matched")
}
