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
	"github.com/dash-app/remote-go/template"
)

// VendorSet - Remote Controller Identifier
type VendorSet struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

// Remote - Remote set (aircon, light etc...)
type Remote struct {
	aircon map[VendorSet]aircon.Remote
}

// Init - Initialize remote controller
func Init() *Remote {
	return &Remote{
		aircon: map[VendorSet]aircon.Remote{
			{Vendor: "daikin", Model: "daikin01"}:         daikin01.New(),
			{Vendor: "daikin", Model: "daikin02"}:         daikin02.New(),
			{Vendor: "daikin", Model: "daikin03"}:         daikin03.New(),
			{Vendor: "daikin", Model: "daikin04"}:         daikin04.New(),
			{Vendor: "fujitsu", Model: "fujitsu01"}:       fujitsu01.New(),
			{Vendor: "mitsubishi", Model: "mitsubishi02"}: mitsubishi02.New(),
			{Vendor: "panasonic", Model: "panasonic01"}:   panasonic01.New(),
		},
	}
}

func (r *Remote) GetTemplate(kind, vendor, model string) (*template.Template, error) {
	switch kind {
	case "AIRCON":
		e, err := r.GetAircon(vendor, model)
		if err != nil {
			return nil, err
		}
		return e.Template(), nil
	default:
		return nil, errors.New("unsupport kind")
	}
}

func (r *Remote) GetAircon(vendor, model string) (aircon.Remote, error) {
	if ac, ok := r.aircon[VendorSet{Vendor: vendor, Model: model}]; ok {
		return ac, nil
	}
	return nil, errors.New("not found")
}

// AvailableAircons - Get vendor/models name
func (r *Remote) AvailableAircons() map[string][]string {
	var set []VendorSet
	for k := range r.aircon {
		set = append(set, k)
	}
	return extractVendorSet(set)
}

// extractVendorSet - Convert to map[string][]string format
func extractVendorSet(from []VendorSet) map[string][]string {
	result := make(map[string][]string)
	for _, k := range from {
		if result[k.Vendor] == nil {
			result[k.Vendor] = []string{k.Model}
		} else {
			result[k.Vendor] = append(result[k.Vendor], k.Model)
		}
	}

	return result
}
