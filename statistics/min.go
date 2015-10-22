package statistics

import (
	"errors"
	"math"
)

// Min returns the smallest value from a supplied set of numerical values
func Min(numbers ...float64) (float64, error) {

	if len(numbers) == 0 {
		return 0.0, nil
	}

	min := numbers[0]

	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!, Occurred because the supplied number argument cannot be recognised as numeric value")
		}
		if number < min {
			min = number
		}
	}

	return min, nil
}
