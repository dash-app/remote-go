package remote

import "github.com/dash-app/remote-go/template"

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

type Remote struct {
	Vendor   string
	Model    string
	Template *template.Template
}

type IRCode struct {
	Code [][]int
}

//type Remote interface {
//	Generate()
//}

type State interface {
	Save() error
	Load() error
}
