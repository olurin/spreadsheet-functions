package spreadsheet

import "testing"

func TestAbs(t *testing.T) {

	// number:  The real number of which you want the absolute value.
	number := 9.0
	if Abs(number) != 9 {
		t.Errorf("Absolute (Abs) Function Failed ... ")
	}

	if Abs(-9) != 9 {
		t.Errorf("Absolute (Abs) Function Failed ... ")
	}
}
