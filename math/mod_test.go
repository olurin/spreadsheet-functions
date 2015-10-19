package mathlib

import "testing"

func TestMod(t *testing.T) {

	result, err := Mod(3, 2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != 1 {
		t.Errorf("Mod Function Failed Mod(3, 2) ... ")
	}

	result, err = Mod(-3, 2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != 1 {
		t.Errorf("Mod Function Failed  Mod(-3, 2) ... ")
	}

	result, err = Mod(3, -2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != -1 {
		t.Errorf("Mod Function Failed  Mod(3, -2) ... ")
	}

	result, err = Mod(-3, -2)

	if err != nil {
		t.Errorf("Mod Function returned an error")
	}

	if result != -1 {
		t.Errorf("Mod Function Failed Mod(-3, -2) ... ")
	}

}
