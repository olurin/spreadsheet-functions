package mathlib

import "testing"

func TestLN(t *testing.T) {

	x := 0.5
	y := 100.00

	result, err := LN(x)

	if err != nil {
		t.Errorf("LN Function returned an error")
	}

	if result != -0.6931471805599453 {
		t.Errorf("LN Function Failed ... ")
	}

	result, err = LN(y)

	if err != nil {
		t.Errorf("LN Function returned an error")
	}

	if result != 4.605170185988092 {
		t.Errorf("LN Function Failed ... ")
	}

}
