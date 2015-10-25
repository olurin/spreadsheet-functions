package statistics

import (
	"errors"
	"math"
)

// Average returns the arithmetic mean of a list of supplied numbers
func Average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, nil
	}

	var total float64
	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!	-	Arguments are supplied directly to the Product function can not be interpreted as numeric values")
		}
		total = total + number
	}

	return total / float64(len(numbers)), nil
}
