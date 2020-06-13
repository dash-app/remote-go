package remote

import (
	"github.com/dash-app/remote-go/remote/aircon/mitsubishi/sg174"
)

//type remote interface {
//	// Generate Entry from JSON payload
//	Generate(opts map[string]interface{})
//
//	// Convert state for save
//
//	// Get Template
//
//	// Get Code
//}

//type Remote struct {
//	Vendor   string
//	Model    string
//	Sender   func(signal []int) error
//	Template *template.Template
//}

type Remote interface {
	Generate()
}

func New(vendor string, model string) *Remote {
	switch {
	case vendor == "mitsubishi" && model == "sg174":
		sg174.New()
	}

	return nil
}
