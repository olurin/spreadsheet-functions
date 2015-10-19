package mathlib

import "testing"

func TestLog10(t *testing.T) {

	x := 86.0
	y := 100.0
	z := 10.0

	result, err := Log10(x)
	if err != nil {
		t.Errorf("Log10 Function returned an error")
	}

	if result != 1.9344984512435675 {
		t.Errorf("Log10 Function Failed ... ")
	}

	result, err = Log10(y)

	if err != nil {
		t.Errorf("Log10 Function returned an error")
	}

	if result != 2 {
		t.Errorf("Log10 Function Failed ... ")
	}

	result, err = Log10(z)

	if err != nil {
		t.Errorf("Log10 Function returned an error")
	}

	if result != 1 {
		t.Errorf("Log10 Function Failed ... ")
	}
}
