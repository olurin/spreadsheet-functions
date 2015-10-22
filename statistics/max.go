package statistics

import (
	"errors"
	"math"
)

// Max returns the largest value from a supplied set of numerical values
func Max(numbers ...float64) (float64, error) {

	if len(numbers) == 0 {
		return 0.0, nil
	}

	max := numbers[0]

	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!, Occurred because the supplied number argument cannot be recognised as numeric value")
		}
		if number > max {
			max = number
		}
	}

	return max, nil
}
