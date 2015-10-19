package mathlib

import (
	"errors"
	"math"
)

// The ABS function returns the absolute value
// (ie. the modulus) of any supplied number

// The Syntax of the function is: ABS(number)
//  ABS(number)

func abs(number float64) float64 {
	switch {
	case number < 0:
		return -number
	case number == 0:
		return 0
	}
	return number
}

// Abs function returns the absolute value of a number.
// The absolute value of a number is the number without its sign.
func Abs(number float64) (float64, error) {

	// Validate Number
	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!, Occurs if the supplied number argument cannot be recognised as numeric value")
	}

	return abs(number), nil
}
