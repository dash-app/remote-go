package remotego

import (
	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/aircon/daikin/daikin01"
	"github.com/dash-app/remote-go/aircon/daikin/daikin02"
	"github.com/dash-app/remote-go/aircon/daikin/daikin03"
	"github.com/dash-app/remote-go/aircon/daikin/daikin04"
	"github.com/dash-app/remote-go/aircon/fujitsu/fujitsu01"
	"github.com/dash-app/remote-go/aircon/mitsubishi/mitsubishi02"
	"github.com/dash-app/remote-go/aircon/panasonic/panasonic01"
)

// VendorSet - Remote Controller Identifier
type VendorSet struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

// Remote - Remote set (aircon, light etc...)
type Remote struct {
	Aircon map[VendorSet]aircon.Remote
}

type RemoteImpl interface {
	GetVendorAndModels() map[string][]string
}

// Init - Initialize remote controller
func Init() *Remote {
	return &Remote{
		Aircon: map[VendorSet]aircon.Remote{
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

// AvailableModels - Get vendor/models name
func (r *Remote) AvailableModels() map[string][]string {
	result := make(map[string][]string)
	for k := range ac {
		if result[k.Vendor] == nil {
			result[k.Vendor] = []string{k.Model}
		} else {
			result[k.Vendor] = append(result[k.Vendor], k.Model)
		}
	}

	return result
}
