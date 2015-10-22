package formulas

import "github.com/TaperBox/formulas/statistics"

// Max returns the largest value from a supplied set of numerical values
func Max(numbers ...float64) (float64, error) { return statistics.Max(numbers...) }
