package formulas

import "github.com/TaperBox/formulas/statistics"

// Max returns the largest value from a supplied set of numerical values
func Max(numbers ...float64) (float64, error) { return statistics.Max(numbers...) }

// Min returns the smallest value from a supplied set of numerical values
func Min(numbers ...float64) (float64, error) { return statistics.Min(numbers...) }

//Large returns the Kth LARGEST value from a list of supplied numbers, for a given value K
func Large(index int, numbers ...float64) (float64, error) { return statistics.Large(index, numbers...) }

//Small returns the Kth SMALLEST value from a list of supplied numbers, for a given value K
func Small(index int, numbers ...float64) (float64, error) { return statistics.Small(index, numbers...) }

// Average returns the arithmetic mean of a list of supplied numbers
func Average(numbers ...float64) (float64, error) { return statistics.Average(numbers...) }

// Median returns the statistical median (the middle value) of a list of supplied numbers
func Median(numbers ...float64) (float64, error) { return statistics.Median(numbers...) }

// Geomean returns the geometric mean of a set of supplied numbers
func Geomean(numbers ...float64) (float64, error) { return statistics.Geomean(numbers...) }

// Harmean returns the harmonic mean of a set of supplied numbers
func Harmean(numbers ...float64) (float64, error) { return statistics.Harmean(numbers...) }

// TrimMean returns the mean of the interior of a supplied set of values
func TrimMean(percent float64, arrays ...float64) (float64, error) {
	return statistics.Harmean(numbers...)
}
