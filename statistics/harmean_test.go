package statistics

import "testing"

func TestHarmean(t *testing.T) {

	x, err := Harmean(2.5, 3, 0.5, 1, 3)

	if x != 1.2295081967213115 || err != nil {
		t.Errorf("Harmean Function Failed ... ")
	}
}
