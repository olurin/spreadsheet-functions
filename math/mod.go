package mathlib

import (
	"math"

	errs "github.com/TaperBox/formulas/errors"
)

// Mod returns the remainder after number is divided by divisor. The result has the same sign as divisor.
func Mod(number, divisor float64) (float64, errs.ErrorType) {
	// Error
	err := errs.ErrorType{}
	if math.IsNaN(divisor) || math.IsNaN(number) {
		err.Type = errs.ErrorCode(6) // #NUM!
		err.Formula = "Mod"
		err.Error = "Occurs if any of the supplied arguments are non-numeric."
		return 0.0, err
	}

	if divisor == 0 {
		err.Type = errs.ErrorCode(2) // #DIV/0!
		err.Formula = "Mod"
		err.Error = "Occurs if the supplied denominator argument is zero"
		return 0.0, err
	}

	sign := Sign(divisor)

	n, err := Abs(number)
	if err != (errs.ErrorType{}) {
		return 0.0, err
	}

	d, err := Abs(divisor)
	if err != (errs.ErrorType{}) {
		return 0.0, err
	}

	return sign * (n - (d)*float64(int((n)/(d)))), err
}
