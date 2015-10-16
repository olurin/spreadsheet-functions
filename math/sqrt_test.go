package spreadsheet

import "testing"

func TestSqrt(t *testing.T) {
	if Sqrt(81) != 9 {
		t.Errorf("Sqrt Function Failed ... ")
	}
}
