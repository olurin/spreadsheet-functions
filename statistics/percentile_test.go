package statistics

import (
	"log"
	"testing"
)

func TestPercentile(t *testing.T) {

	nums := []float64{2, 1, 6, 4, 3, 5}

	x, err := Percentile(20, nums...)

	if x != 2 || err != nil {
		t.Errorf("Percentile function failed ... ")
	}

	x, err = Percentile(60, nums...)

	if x != 4 || err != nil {
		t.Errorf("Percentile function failed ... ")
	}

	x, err = Percentile(50, nums...)
	if x != 4 || err != nil {
		log.Println(x)
		t.Errorf("Percentile function failed ... ")
	}

}
