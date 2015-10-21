package mathlib

import (
	"errors"
	"math"
)

// Quotient returns the integer portion of a division between two supplied numbers
func Quotient(numerator, denominator float64) (int, error) {
	if denominator == 0 {
		return 0.0, errors.New("#DIV/0!	- Occurred because the supplied denominator argument is zero")
	}

	if math.IsNaN(numerator) || math.IsNaN(denominator) {
		return 0.0, errors.New("#VALUE!	- Occurred because  the supplied arguments are non-numeric ")
	}

	return int(numerator / denominator), nil
}
