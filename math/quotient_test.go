package spreadsheet

import "testing"

func TestQuotient(t *testing.T) {

	if Quotient(5, 2) != 2 {
		t.Errorf("Quotient Function Failed ... ")
	}
	if Quotient(4.5, 3.1) != 1 {
		t.Errorf("Quotient Function Failed ... ")
	}

	if Quotient(-10, 3) != -3 {
		t.Errorf("Quotient Function Failed ... ")
	}
}
