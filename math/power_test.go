package mathlib

import "testing"

func TestPower(t *testing.T) {

	p, err := Power(10, 4)
	if p != 10000 && err != nil {
		t.Errorf("Product Function Failed ... ")
	}

}
