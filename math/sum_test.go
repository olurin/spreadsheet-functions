package mathlib

import "testing"

func TestSum(t *testing.T) {
	if Sum(12, 34, 65, 34, 23) != 168 {
		t.Errorf("Sum Function Failed ... ")
	}

	if Sum(12, 23) != 35 {
		t.Errorf("Sum Function Failed ... ")
	}

	if Sum(-12, 23) != 11 {
		t.Errorf("Sum Function Failed ... ")
	}
}
