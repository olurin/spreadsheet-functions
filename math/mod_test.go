package mathlib

import (
	"testing"

	errs "github.com/TaperBox/formulas/errors"
)

func TestMod(t *testing.T) {

	result, err := Mod(3, 2)

	if err != (errs.ErrorType{}) {
		t.Errorf("Mod Function returned an error")
	}

	if result != 1 {
		t.Errorf("Mod Function Failed Mod(3, 2) ... ")
	}

	result, err = Mod(-3, 2)

	if err != (errs.ErrorType{}) {
		t.Errorf("Mod Function returned an error")
	}

	if result != 1 {
		t.Errorf("Mod Function Failed  Mod(-3, 2) ... ")
	}

	result, err = Mod(3, -2)

	if err != (errs.ErrorType{}) {
		t.Errorf("Mod Function returned an error")
	}

	if result != -1 {
		t.Errorf("Mod Function Failed  Mod(3, -2) ... ")
	}

	result, err = Mod(-3, -2)

	if err != (errs.ErrorType{}) {
		t.Errorf("Mod Function returned an error")
	}

	if result != -1 {
		t.Errorf("Mod Function Failed Mod(-3, -2) ... ")
	}

}
