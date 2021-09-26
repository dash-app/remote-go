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
	"github.com/dash-app/remote-go/light"
	"github.com/dash-app/remote-go/light/hitachi/ira03h"
	"github.com/dash-app/remote-go/light/odelic/rc701w"
	"github.com/dash-app/remote-go/template"
)

// VendorSet - Remote Controller Identifier
type VendorSet struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

// Remote - Remote set (aircon, light etc...)
type Remote struct {
	aircon map[VendorSet]*RemoteSet
	light  map[VendorSet]*RemoteSet
}

// RemoteSet - Abstract remote structure
type RemoteSet struct {
	aircon aircon.Remote
	light  light.Remote
}

// Init - Initialize remote controller
func Init() *Remote {
	return &Remote{
		aircon: map[VendorSet]*RemoteSet{
			{Vendor: "daikin", Model: "daikin01"}:         {aircon: daikin01.New()},
			{Vendor: "daikin", Model: "daikin02"}:         {aircon: daikin02.New()},
			{Vendor: "daikin", Model: "daikin03"}:         {aircon: daikin03.New()},
			{Vendor: "daikin", Model: "daikin04"}:         {aircon: daikin04.New()},
			{Vendor: "fujitsu", Model: "fujitsu01"}:       {aircon: fujitsu01.New()},
			{Vendor: "mitsubishi", Model: "mitsubishi02"}: {aircon: mitsubishi02.New()},
			{Vendor: "panasonic", Model: "panasonic01"}:   {aircon: panasonic01.New()},
		},
		light: map[VendorSet]*RemoteSet{
			{Vendor: "hitachi", Model: "ir-a03h"}: {light: ira03h.New()},
			{Vendor: "odelic", Model: "rc701w"}:   {light: rc701w.New()},
		},
	}
}

// Aircon - Get Aircon
func (rs *RemoteSet) Aircon() aircon.Remote {
	return rs.aircon
}

// Light - Get light
func (rs *RemoteSet) Light() light.Remote {
	return rs.light
}

// Template - Get template
func (rs *RemoteSet) Template() *template.Template {
	if rs.aircon != nil {
		return rs.aircon.Template()
	}
	if rs.light != nil {
		return rs.light.Template()
	}
	return nil
}

// Get - Get remote-set
func (r *Remote) Get(kind, vendor, model string) (*RemoteSet, error) {
	switch kind {
	case "AIRCON":
		return r.aircon[VendorSet{Vendor: vendor, Model: model}], nil
	case "LIGHT":
		return r.light[VendorSet{Vendor: vendor, Model: model}], nil
	default:
		return nil, errors.New("unsupport kind")
	}
}

// AvailableAircons - Get vendor/models name
func (r *Remote) AvailableAircons() map[string][]string {
	var set []VendorSet
	for k := range r.aircon {
		set = append(set, k)
	}
	return extractVendorSet(set)
}

// AvailableLights - Get vendor/models name
func (r *Remote) AvailableLights() map[string][]string {
	var set []VendorSet
	for k := range r.light {
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
