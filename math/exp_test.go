package spreadsheet

import "testing"

func TestEXP(t *testing.T) {
	if Exp(100) != 2.6881171418161356e+43 {
		t.Errorf("Exp Function Failed ... ")
	}

	if Exp(0.1) != 1.1051709180756477 {
		t.Errorf("Exp Function Failed ... ")
	}

	if Exp(-5) != 0.006737946999085467 {
		t.Errorf("Exp Function Failed ... ")
	}
}
