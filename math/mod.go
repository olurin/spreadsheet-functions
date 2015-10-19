package mathlib

import (
	"errors"
	"math"
)

// Mod returns the remainder after number is divided by divisor. The result has the same sign as divisor.
func Mod(number, divisor float64) (float64, error) {

	if math.IsNaN(divisor) && math.IsNaN(number) {
		return 0.0, errors.New("#NUM Error! - supplied number argument is negative or zero")
	}

	if divisor == 0 {
		return 0.0, errors.New("#DIV/0! error value")
	}

	sign := Sign(divisor)

	n, err := Abs(number)
	if err != nil {
		return 0.0, err
	}

	d, err := Abs(divisor)
	if err != nil {
		return 0.0, err
	}

	return sign * (n - (d)*float64(int((n)/(d)))), nil
}
