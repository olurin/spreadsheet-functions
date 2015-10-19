package mathlib

import (
	"errors"
	"math"
)

// Log10 Returns the base-10 logarithm of a number.
func Log10(number float64) (float64, error) {

	// Validate Number
	if number <= 0 {
		return 0.0, errors.New("#NUM Error! - supplied number argument is negative or zero")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE Error! - supplied number argument is not a numeric value")
	}

	return math.Log10(number), nil
}
