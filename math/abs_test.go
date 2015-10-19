package mathlib

import "testing"

func TestAbs(t *testing.T) {

	x := 9.0
	y := -10.0
	z := -10.34

	t1, err := Abs(x)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 9 {
		t.Errorf("Abs Function failed")
	}

	t1, err = Abs(y)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 10.0 {
		t.Errorf("Abs Function failed")
	}

	t1, err = Abs(z)
	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if t1 != 10.34 {
		t.Errorf("Abs Function failed")
	}
}
