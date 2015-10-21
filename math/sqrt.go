package mathlib

import (
	"errors"
	"math"
)

// Sqrt returns the positive square root of a given number
func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0.0, errors.New("#NUM!	-	Occurred because the supplied number argument is negative")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Sqrt(number), nil
}
