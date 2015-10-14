package spreadsheet

import "testing"

func TestTrigonometry(t *testing.T) {
	if PI() != 3.14159265358979 {
		t.Errorf("PI function failed ... ")
	}

	if SqrtPI(1) != 1.7724538509055152 {
		t.Errorf("SqrtPI function failed ... ")
	}

	if SqrtPI(2) != 2.5066282746309994 {
		t.Errorf("SqrtPI function failed ... ")
	}

	if Degree(PI()) != 180 {
		t.Errorf("Degree function failed ... ")
	}

	if Radians(270) != 4.712388980384685 {
		t.Errorf("Radians function failed ... ")
	}
}
