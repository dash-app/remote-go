package test

import (
	"fmt"
	"reflect"

	"github.com/dash-app/remote-go/appliances"
	"github.com/dash-app/remote-go/hex"
)

// Original - Original Code
type Original struct {
	Code [][]int
}

// RemoteTest - Test object for Remote
type RemoteTest struct {
	Title    string
	Remote   appliances.RemoteController
	Original []*Original
	Request  *appliances.Request
}

// Compare - validate both codes
func (rt *RemoteTest) Compare() error {
	code, err := rt.Remote.Generate(rt.Request)
	if err != nil {
		return err
	}

	// Output both entries
	for i, c := range code {
		if ok := reflect.DeepEqual(rt.Original[i].Code, c.Code); !ok {
			errs := "\n"
			errs += fmt.Sprintln("-- Expected:")
			for _, v := range hex.Format(rt.Original[i].Code) {
				errs += fmt.Sprintf("\t%s", v)
			}
			errs += fmt.Sprintln("-- Generated:")
			for _, v := range hex.Format(c.Code) {
				errs += fmt.Sprintf("\t%s", v)
			}

			return fmt.Errorf("Compare failed!! (step: %d) %v", i, errs)
		}
	}

	return nil
}
