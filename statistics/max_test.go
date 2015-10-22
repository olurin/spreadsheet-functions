package statistics

import "testing"

func TestMax(t *testing.T) {

	max, err := Max(4, 34, 1, 5, 3, -4, 3)

	if err != nil || max != 34 {
		t.Errorf("LCM Function returned an error")
	}
}
