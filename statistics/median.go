package statistics

import "sort"

// Median returns the statistical median (the middle value) of a list of supplied numbers
func Median(numbers ...float64) (float64, error) {
	l := len(numbers)

	if l <= 0 {
		return 0.0, nil
	}
	sort.Sort(Nums(numbers))
	if l%2 == 0 {
		s := numbers[(l-1)/2] + numbers[(l+1)/2]
		return s / 2, nil
	}

	return numbers[l/2], nil
}
