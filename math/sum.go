package mathlib

import (
	"errors"
	"math"
)

// Sum returns the sum of a supplied list of numbers
func Sum(numbers ...float64) (float64, error) {
	var sum float64
	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!	-	Arguments are supplied directly to the Product function can not be interpreted as numeric values")
		}
		sum = sum + number
	}

	return sum, nil
}
