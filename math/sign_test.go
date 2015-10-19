package mathlib

import "testing"

func TestSign(t *testing.T) {

	if Sign(9043) != 1 {
		t.Errorf("Sign Function Failed ... ")
	}

	if Sign(-9043) != -1 {
		t.Errorf("Sign Function Failed ... ")
	}

	if Sign(1) != 1 {
		t.Errorf("Sign Function Failed ... ")
	}

}
