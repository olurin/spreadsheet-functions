package mathlib

import "testing"

func TestLCM(t *testing.T) {

	a := []int{1, 5}
	b := []int{15, 10, 25}
	c := []int{1, 8, 12}
	x := []int{7, 2}

	result, err := LCM(a...)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 5 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(b...)

	if err != nil {
		t.Errorf("Abs Function returned an error")
	}

	if result != 150 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(c...)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 24 {
		t.Errorf("LCM Function Failed ")
	}

	result, err = LCM(x...)

	if err != nil {
		t.Errorf("LCM Function returned an error")
	}

	if result != 14 {
		t.Errorf("LCM Function Failed ")
	}

}
