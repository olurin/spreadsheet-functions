package mathlib

import "testing"

func TestSqrt(t *testing.T) {
	x, err := Sqrt(81)

	if x != 9 && err != nil {
		t.Errorf("Sqrt Function Failed ... ")
	}

	x, err = Sqrt(-81)

	if x != 0.0 && err == nil {
		t.Errorf("Sqrt Function Failed ... ")
	}
}
