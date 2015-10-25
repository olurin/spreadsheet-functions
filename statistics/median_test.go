package statistics

import "testing"

func TestMedian(t *testing.T) {

	x, err := Median(13, 18, 13, 14, 13, 16, 14, 21, 13)

	if x != 14 && err != nil {
		t.Errorf("Median Function Failed ... ")
	}

	x, err = Median(1, 2, 3, 4)

	if x != 2.5 && err != nil {
		t.Errorf("Median Function Failed ... ")
	}

}
