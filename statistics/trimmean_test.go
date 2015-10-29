package statistics

import (
	"testing"

	"github.com/TaperBox/formulas/math"
)

func TestTrimMean(t *testing.T) {
	x, err := TrimMean(15, 2, 4, 6, 7, 11, 21, 81, 90, 105, 121)

	if mathlib.Round(x, 3) != 44.80 || err != nil {
		t.Errorf("TrimMean Function Failed ... ")
	}

	x1 := []float64{0.5, 6, 7, 7, 8, 8, 9, 9, 16, 24}

	x, err = TrimMean(10, x1...)
	if mathlib.Round(x, 3) != 9.45 || err != nil {
		t.Errorf("TrimMean Function Failed ... ")
	}

	x, err = TrimMean(20, x1...)
	if mathlib.Round(x, 3) != 8.75 || err != nil {
		t.Errorf("TrimMean Function Failed ... ")
	}

	x, err = TrimMean(40, x1...)
	if mathlib.Round(x, 3) != 8.00 || err != nil {
		t.Errorf("TrimMean Function Failed ... ")
	}
}
