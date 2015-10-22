package mathlib

import "testing"

func TestProduct(t *testing.T) {

	y := []float64{100, 10, 45}

	p, err := Product(45, 2, 23)
	if p != 2070 && err != nil {
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Product(y...)
	if p != 45000 && err != nil {
		t.Errorf("Product Function Failed ... ")
	}
}
