package mathlib

import "testing"

func TestMod(t *testing.T) {
	if Mod(3, 2) != 1 {
		t.Errorf("Mod Function Failed Mod(3, 2) ... ")
	}

	if Mod(-3, 2) != 1 {
		t.Errorf("Mod Function Failed  Mod(-3, 2) ... ")
	}

	if Mod(3, -2) != -1 {
		t.Errorf("Mod Function Failed  Mod(3, -2) ... ")
	}

	if Mod(-3, -2) != -1 {
		t.Errorf("Mod Function Failed Mod(-3, -2) ... ")
	}

}
