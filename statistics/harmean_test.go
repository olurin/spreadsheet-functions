package statistics

import "testing"

func TestHarmean(t *testing.T) {

	x, err := Harmean(3, 0.5, 1, 3)

	if x != 2.5 && err != nil {
		t.Errorf("Harmean Function Failed ... ")
	}
}
