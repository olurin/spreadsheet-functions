package mathlib

import (
	"errors"
	"math"
)

// Power function returns the result of a given number raised to a supplied power
func Power(number, power float64) (float64, error) {
	if math.IsNaN(number) || math.IsNaN(power) {
		return 0.0, errors.New("#VALUE Error! - supplied arguments are non-numeric")
	}

	return math.Pow(number, power), nil
}
