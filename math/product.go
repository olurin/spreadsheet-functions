package mathlib

import (
	"errors"
	"math"
)

// Product returns the product of a supplied list of numbers
func Product(numbers []float64) (float64, error) {

	product := 1.0

	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!	-	Arguments are supplied directly to the Product function can not be interpreted as numeric values")
		}
		product = product * number
	}
	return product, nil
}
