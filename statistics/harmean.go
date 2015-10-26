package statistics

import (
	"errors"
	"math"
)

// Harmean returns the harmonic mean of a set of supplied numbers
func Harmean(numbers ...float64) (float64, error) {
	var h float64

	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!	-	Arguments are supplied directly to the Product function can not be interpreted as numeric values")
		}

		if number < 0 {
			return 0.0, errors.New("#NUM!	-	Occurred because some of the supplied numeric values are negative")
		}
		h += (1 / number)
	}

	return float64(len(numbers)) / h, nil
}
