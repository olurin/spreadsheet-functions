package mathlib

import "testing"

func TestEXP(t *testing.T) {

	x := 100.0
	y := 0.1
	z := -5.0

	t1, err := Exp(x)

	if err != nil {
		t.Errorf("Exp Function returned an error")
	}

	if t1 != 2.6881171418161356e+43 {
		t.Errorf("Exp Function Failed ... ")
	}

	t1, err = Exp(y)

	if err != nil {
		t.Errorf("Exp Function returned an error")
	}

	if t1 != 1.1051709180756477 {
		t.Errorf("Exp Function Failed ... ")
	}

	t1, err = Exp(z)

	if err != nil {
		t.Errorf("Exp Function returned an error")
	}

	if t1 != 0.006737946999085467 {
		t.Errorf("Exp Function Failed ... ")
	}
}
