package statistics

import (
	"errors"
	"math"
	"sort"
)

// TrimMean returns the mean of the interior of a supplied set of values
// Array  -	An array of numeric values, for which you want to calculate the trimmed mean
// Percent	-	The percentage of values that you want to be discarded from the supplied array
func TrimMean(percent float64, arrays ...float64) (float64, error) {

	// First find n = number of observations
	n := float64(len(arrays))

	// Reorder them as "order statistics" Xi from the smallest to the largest
	sort.Sort(SmallNums(arrays))

	if math.IsNaN(percent) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the percent argument cannot be interpreted as a Numeric value")
	}

	// Find lower case p=P/100 = proportion trimmed
	p := percent / 100.00

	if p < 0 || p > 1 {
		return 0.0, errors.New("#NUM!, Occurred the supplied percent argument is < 0 or > 1")
	}

	// If np is an integer use k=np and trim k observations at both ends.
	k := math.Floor((n * p) / 2)

	//r = remaining observations = n−2k
	r := n - (2 * k)

	k = math.Floor(k)

	// Trimmed mean = (1/R)(Xk+1+Xk+2+…+Xn−k)
	sum := 0.0
	for i := int(k); i <= int(n-k-1); i++ {
		sum += arrays[i]
	}
	return (1 / r) * sum, nil
}
