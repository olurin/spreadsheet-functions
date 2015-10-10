package spreadsheet

import "testing"

func TestGCD(t *testing.T) {
	if GCD(203, 91, 77) != 7 {
		t.Errorf("GCD Function Failed ")
	}

	if GCD(77, 91, 203) != 7 {
		t.Errorf("GCD Function Failed ")
	}

	if GCD(5, 2) != 1 {
		t.Errorf("GCD Function Failed ")
	}

	if GCD(24, 36) != 12 {
		t.Errorf("GCD Function Failed ")
	}

	if GCD(5, 0) != 5 {
		t.Errorf("GCD Function Failed ")
	}

	if GCD(15, 10, 25) != 5 {
		t.Errorf("GCD Function Failed ")
	}

}
