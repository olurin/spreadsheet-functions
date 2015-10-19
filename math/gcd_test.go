package mathlib

import "testing"

func TestGCD(t *testing.T) {

	a := []int{203, 91, 77}
	b := []int{77, 91, 203}
	c := []int{5, 2}
	x := []int{24, 36}
	y := []int{5, 0}
	z := []int{15, 10, 25}

	result, err := GCD(a)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 7 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(b)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 7 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(c)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 1 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(x)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 12 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(y)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 5 {
		t.Errorf("GCD Function Failed ")
	}

	result, err = GCD(z)

	if err != nil {
		t.Errorf("GCD Function returned an error")
	}

	if result != 5 {
		t.Errorf("GCD Function Failed ")
	}

}
