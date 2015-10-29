package statistics

import "testing"

func TestSmall(t *testing.T) {

	large, err := Large(1, 4, 34, 1, 5, 3, -4, 3)

	if large != 34 || err != nil {
		t.Errorf("Large Function returned an error")
	}

	large, err = Large(10, 4, 34, 1, 5, 3, -4, 3)

	if large != 0.0 || err == nil {
		t.Errorf("Large Function returned an error")
	}
}
