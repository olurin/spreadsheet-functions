package mathlib

import "math"

// Round rounds a number up or down, to a given number of digits
// The syntax of the function is:
// ROUND( number, numDigits )
func Round(number float64, numDigits int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(numDigits))
	intermed := number * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
