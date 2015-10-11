package spreadsheet

import "testing"

func TestLN(t *testing.T) {
	if LN(0.5) != -0.6931471805599453 {
		t.Errorf("LN Function Failed ... ")
	}

	if LN(100) != 4.605170185988092 {
		t.Errorf("LN Function Failed ... ")
	}
}
