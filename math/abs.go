package mathlib

import (
	"math"

	errs "github.com/TaperBox/formulas/errors"
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
func Abs(number float64) (float64, errs.ErrorType) {
	// Error
	err := errs.ErrorType{}
	// Validate Number
	if math.IsNaN(number) {
		// If you get an error from the Excel Abs function, this is likely to be the #VALUE! error
		err.Type = errs.ErrorCode(3)
		err.Formula = "Abs"
		err.Error = "Occurs if the supplied number argument cannot be recognised as numeric value"
		return 0.0, err
	}

	return abs(number), err
}
