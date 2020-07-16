package test

import (
	"fmt"
	"reflect"

	"github.com/dash-app/remote-go/aircon"
	"github.com/dash-app/remote-go/hex"
	"github.com/dash-app/remote-go/remote"
)

type ACOriginals struct {
	Code [][]int
}

type ACTestEntry struct {
	Title    string
	Remote   remote.Aircon
	Original []*ACOriginals
	Entry    *aircon.Entry
}

// Validate - valiate entry by template validator
// TODO: Remove from here (should be move to entry?)
func (te *ACTestEntry) Validate() error {
	entry := te.Entry
	template := te.Remote.Template()
	template.Aircon.Operation.Validate(entry.Operation)

	if t := template.Aircon.Modes[entry.Mode].Temp; t != nil {
		if err := t.Validate(entry.Temp); err != nil {
			return err
		}
	}

	if t := template.Aircon.Modes[entry.Mode].Fan; t != nil {
		if err := t.Validate(entry.Fan); err != nil {
			return err
		}
	}

	if t := template.Aircon.Modes[entry.Mode].HorizontalVane; t != nil {
		if err := t.Validate(entry.HorizontalVane); err != nil {
			return err
		}
	}

	if t := template.Aircon.Modes[entry.Mode].VerticalVane; t != nil {
		if err := t.Validate(entry.VerticalVane); err != nil {
			return err
		}
	}

	return nil
}

func (te *ACTestEntry) Compare() error {
	code, err := te.Remote.Generate(te.Entry)
	if err != nil {
		return err
	}

	// Output both entries
	for i, c := range code {
		if ok := reflect.DeepEqual(te.Original[i].Code, c.Code); !ok {
			errs := "\n"
			errs += fmt.Sprintln("-- Expected:")
			for _, v := range hex.Format(te.Original[i].Code) {
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
