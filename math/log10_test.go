package mathlib

import "testing"

func TestLog10(t *testing.T) {

	if Log10(86) != 1.9344984512435675 {
		t.Errorf("Log10 Function Failed ... ")
	}

	if Log10(100) != 2 {
		t.Errorf("Log10 Function Failed ... ")
	}

	if Log10(10) != 1 {
		t.Errorf("Log10 Function Failed ... ")
	}
}
