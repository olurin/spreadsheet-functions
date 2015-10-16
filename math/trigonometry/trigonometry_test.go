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

	if Cos(0.785398163) != 0.707106781467586 {
		t.Errorf("Cos function failed ... ")
	}

	if Cos(0) != 1 {
		t.Errorf("Cos function failed ... ")
	}

	if Cos(-PI()/3) != 0.5000000000000009 {
		t.Errorf("Cos function failed ... ")
	}

	if Cos(Radians(-30)) != 0.8660254037844389 {
		t.Errorf("Cos function failed ... ")
	}
}
