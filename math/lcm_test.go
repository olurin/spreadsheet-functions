package spreadsheet

import "testing"

func TestLCM(t *testing.T) {

	if LCM(1, 5) != 5 {
		t.Errorf("LCM Function Failed ")
	}

	if LCM(15, 10, 25) != 150 {
		t.Errorf("LCM Function Failed ")
	}

	if LCM(1, 8, 12) != 24 {
		t.Errorf("LCM Function Failed ")
	}

	if LCM(7, 2) != 14 {
		t.Errorf("LCM Function Failed ")
	}

}
