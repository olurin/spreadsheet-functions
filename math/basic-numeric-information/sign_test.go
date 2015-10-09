package spreadsheet

import "testing"

func TestSign(t *testing.T) {

	if Sign(9043) != 1 {
		t.Errorf("Sign Function Failed ... ")
	}

	if Sign(-9043) != -1 {
		t.Errorf("Sign Function Failed ... ")
	}

	if Sign(0) != 0 {
		t.Errorf("Sign Function Failed ... ")
	}

}
