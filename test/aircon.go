package test

import (
	"fmt"
	"reflect"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/hex"
	"github.com/dash-app/remote-go/remote"
)

type ACTestEntry struct {
	Title    string
	Remote   remote.Aircon
	Original [][]int
	Entry    *aircon.Entry
}

func (te *ACTestEntry) Compare() error {
	code, err := te.Remote.Generate(te.Entry)
	if err != nil {
		return err
	}

	// Output both entries
	if ok := reflect.DeepEqual(te.Original, code[0].Code); !ok {
		errs := "\n"
		errs += fmt.Sprintln("-- Expected:")
		for _, v := range hex.Format(te.Original) {
			errs += fmt.Sprintf("\t%s", v)
		}
		errs += fmt.Sprintln("-- Generated:")
		for _, v := range hex.Format(code[0].Code) {
			errs += fmt.Sprintf("\t%s", v)
		}

		return fmt.Errorf("Compare failed!! %v", errs)
	}
	return nil
}
