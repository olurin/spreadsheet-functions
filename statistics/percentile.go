package statistics

import (
	"errors"
	"math"
	"sort"

	"github.com/TaperBox/formulas/math"
)

// array	-	The array of data values for which you want to calculate the k'th percentile
// k	-	The value (between 0 and 1) of the required percentile

// Percentile function returns the k'th percentile of a supplied
// range of values for a given value of k
func Percentile(k float64, arrays ...float64) (float64, error) {

	// Check if k is a Non Numeric
	if math.IsNaN(k) {
		return 0.0, errors.New("#VALUE!	-	Occurred because  the supplied value of k is non-numeric")
	}

	if len(arrays) == 0 || (k < 0 || k > 100) {
		return 0.0, errors.New("#NUM!	-	Occurred because the supplied value of k is less than 0 or greater than 100 or the array is empty")
	}

	// Reorder them from the smallest to the largest
	sort.Sort(SmallNums(arrays))

	// n =  \left \lceil \frac{P}{100} \times N  \right \rceil

	n := (k / 100) * float64(len(arrays))

	mathlib.Round(n, 0)

	return arrays[int(n)], nil
}
