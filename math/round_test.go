package mathlib

import "testing"

func TestRound(t *testing.T) {

	if Round(4.05, 0) != 4 {
		t.Errorf("Round function in Math library failed...")
	}
	if Round(4.05, 2) != 4.05 {
		t.Errorf("Round function in Math library failed...")
	}
	if Round(4.05, 1) != 4.1 {
		t.Errorf("Round function in Math library failed...")
	}

	if Round(4.05, 10) != 4.05 {
		t.Errorf("Round function in Math library failed...")
	}
}
