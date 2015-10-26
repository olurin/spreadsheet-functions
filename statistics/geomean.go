package statistics

import (
	"errors"
	"math"
)

// Geomean returns the geometric mean of a set of supplied numbers
func Geomean(numbers ...float64) (float64, error) {
	g := 1.0

	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New("#VALUE!	-	Arguments are supplied directly to the Product function can not be interpreted as numeric values")
		}

		if number < 0 {
			return 0.0, errors.New("#NUM!	-	Occurred because some of the supplied numeric values are negative")
		}
		g *= number
	}
	return math.Pow(g, 1/float64(len(numbers))), nil
}
