package mathlib

import "errors"

func gcd(m, n int) int {
	var r int
	if (m == 0) && (n == 0) {
		return -1
	} else if (m < 0) || (n < 0) {
		return -1
	}

	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	for {
		r = m % n
		if r == 0 {
			break
		}
		m = n
		n = r
	}
	return n
}

// GCD function returns the greatest common divisor of two or more supplied integers.
// The format of the function is :   GCD( number1, number2, ...)
// Where the number arguments are up to 255 numerical values for which you want to calculate the greatest common divisor.
// If any of the supplied numbers are not integers, these values are truncated to integers.
func GCD(numbers []int) (int, error) {

	// Validate Numbers
	if len(numbers) <= 1 {
		return 0, errors.New("Number1, number2, ...    Number1 is required, subsequent numbers are optional. 1 to 255 values. If any value is not an integer, it is truncated.")
	}

	var gcdValue int
	var index = 0

	l := len(numbers) - 1
	for l >= 0 {
		if index == 0 {
			gcdValue = gcd(numbers[l-1], numbers[l])
			l = l - 1
			index = index + 1
			continue
		}
		gcdValue = gcd(numbers[l], gcdValue)
		l = l - 1
	}

	return gcdValue, nil
}
