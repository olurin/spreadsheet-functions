package statistics

import "testing"

func TestMin(t *testing.T) {

	min, err := Min(4, 34, 1, 5, 3, -4, 3)

	if err != nil || min != -4 {
		t.Errorf("LCM Function returned an error")
	}
}
