package mathlib

import "testing"

func TestSum(t *testing.T) {

	x, err := Sum(12, 34, 65, 34, 23)

	if x != 168 && err != nil {
		t.Errorf("Sum Function Failed ... ")
	}

	x, err = Sum(12, 23)

	if x != 35 && err != nil {
		t.Errorf("Sum Function Failed ... ")
	}
	x, err = Sum(-12, 23)
	if x != 11 && err != nil {
		t.Errorf("Sum Function Failed ... ")
	}
}
