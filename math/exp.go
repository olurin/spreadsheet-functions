package mathlib

import (
	"errors"
	"math"
)

// Exp function returns e raised to a given power
func Exp(number float64) (float64, error) {

	// Validate Number
	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!, Occurs if the supplied number argument cannot be recognised as numeric value")
	}

	return math.Exp(number), nil
}
