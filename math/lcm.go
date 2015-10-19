package mathlib

import "errors"

// LCM function
func LCM(numbers []int) (int, error) {

	// Validate Numbers
	if len(numbers) <= 1 {
		return 0, errors.New("Number1, number2, ...    Number1 is required, subsequent numbers are optional. 1 to 255 values. If any value is not an integer, it is truncated.")
	}

	var lcmvalue int
	var index = 0

	l := len(numbers) - 1

	for l >= 0 {

		if index == 0 {
			lcmvalue = (numbers[l-1] * numbers[l]) / gcd(numbers[l-1], numbers[l])
			l = l - 1
			index = index + 1
			continue
		}

		lcmvalue = (numbers[l] * lcmvalue) / gcd(numbers[l], lcmvalue)
		l = l - 1
	}
	return lcmvalue, nil
}
