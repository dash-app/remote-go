package appliances

import (
	"encoding/json"
	"fmt"

	"github.com/dash-app/remote-go/hex"
)

// RemoteController - RemoteController Interface
type RemoteController interface {
	Generate(*Request) ([]*hex.HexCode, error)
	Template() *Template
}

// UpdateResult - Result of Update
type UpdateResult struct {
	Aircon *Request
	Light  *LightState
}

// Kind - Appliance kind
type Kind int32

const (
	// UNKNOWN - Unused or Unknown kind
	UNKNOWN Kind = iota

	// AIRCON - Air Conditioner
	AIRCON

	// LIGHT - Light
	LIGHT

	// SWITCHBOT - SwitchBot
	SWITCHBOT
)

func (k Kind) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}

func (k *Kind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("data should be a string, got %s", data)
	}

	*k = KindFromString(s)

	return nil
}

func KindFromString(s string) Kind {
	switch s {
	case "AIRCON":
		return AIRCON
	case "LIGHT":
		return LIGHT
	case "SWITCHBOT":
		return SWITCHBOT
	case "UNKNOWN":
		fallthrough
	default:
		return UNKNOWN
	}
}

func (k Kind) String() string {
	switch k {
	case AIRCON:
		return "AIRCON"
	case LIGHT:
		return "LIGHT"
	case SWITCHBOT:
		return "SWITCHBOT"
	case UNKNOWN:
		fallthrough
	default:
		return "UNKNOWN"
	}
}
