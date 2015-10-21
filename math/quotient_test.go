package mathlib

import "testing"

func TestQuotient(t *testing.T) {

	result, err := Quotient(5, 2)

	if result != 2 || err != nil {
		t.Errorf("Quotient Function Failed ... ")
	}

	result, err = Quotient(4.5, 3.1)

	if result != 1 || err != nil {
		t.Errorf("Quotient Function Failed ... ")
	}

	result, err = Quotient(-10, 3)

	if result != -3 || err != nil {
		t.Errorf("Quotient Function Failed ... ")
	}
}
