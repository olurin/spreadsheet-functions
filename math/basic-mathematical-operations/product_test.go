package spreadsheet

import "testing"

func TestProduct(t *testing.T) {

	p, err := Product(45, 2, 23)
	if p != 2070 && err != nil {
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Product(100, 10, 45)
	if p != 45000 && err != nil {
		t.Errorf("Product Function Failed ... ")
	}
}
