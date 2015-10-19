package mathlib

// LN(number)
// The LN function syntax has the following arguments:
// Number    Required. The positive real number for which you want the natural logarithm.
// LN is the inverse of the EXP function.

import (
	"errors"
	"math"
)

// LN returns the natural logarithm of a number. Natural logarithms are based on the constant e (2.71828182845904)
func LN(number float64) (float64, error) {
	// Validate Number
	if number <= 0 {
		return 0.0, errors.New("#NUM Error! - supplied number argument is negative or zero")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE Error! - supplied number argument is not a numeric value")
	}

	return math.Log(number), nil
}
