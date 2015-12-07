package formulas

import (
	"errors"
	"math"
)

// ABS function returns the absolute value of a number
// The absolute value of a number is the number without its sign
func ABS(number float64) (float64, error) {
	// Validate Number
	if math.IsNaN(number) {
		return 0.0, errors.New(errorCode(3))
	}
	return abs(number), nil
}

// Exp function returns e raised to a given power
func Exp(number float64) (float64, error) {

	// Validate Number
	if math.IsNaN(number) {
		return 0.0, errors.New(errorCode(3))
	}
	return math.Exp(number), nil
}

// GCD function returns the greatest common divisor of two or more supplied integers.
// The format of the function is :   GCD( number1, number2, ...)
// Where the number arguments are up to 255 numerical values for which you want to calculate the greatest common divisor.
// If any of the supplied numbers are not integers, these values are truncated to integers.
func GCD(numbers ...int) (int, error) {

	// Validate Numbers
	if len(numbers) <= 1 {
		return 0, errors.New(errorCode(6))
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

// LCM function returns the least common multiple of two or more supplied integers
// The format of the function is LCM( number1, number2, ...)
// Where the number arguments are up to 255 numerical values for which you want to calculate the least common multiple.
func LCM(numbers ...int) (int, error) {

	// Validate Numbers
	if len(numbers) <= 1 {
		return 0, errors.New(errorCode(6))
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

// LN returns the natural logarithm of a number. Natural
// logarithms are based on the constant e (2.71828182845904)
func LN(number float64) (float64, error) {
	// Validate Number
	if number <= 0 {
		return 0.0, errors.New(errorCode(6))
	}

	if math.IsNaN(number) {
		return 0.0, errors.New(errorCode(3))
	}

	return math.Log(number), nil
}

// Log10 Returns the base-10 logarithm of a number.
func Log10(number float64) (float64, error) {

	// Validate Number
	if number <= 0 {
		return 0.0, errors.New(errorCode(6))
	}

	if math.IsNaN(number) {
		return 0.0, errors.New(errorCode(3))
	}

	return math.Log10(number), nil
}

// Mod returns the remainder after number is divided by divisor. The result has the same sign as divisor.
func Mod(number, divisor float64) (float64, error) {

	if math.IsNaN(divisor) || math.IsNaN(number) {
		return 0.0, errors.New(errorCode(6))
	}

	if divisor == 0 {
		return 0.0, errors.New(errorCode(2))
	}

	sign := Sign(divisor)

	n := abs(number)
	d := abs(divisor)

	return sign * (n - (d)*float64(int((n)/(d)))), nil
}

// Sign function returns the sign (+1, -1 or 0) of a supplied number. ie.
// if the number is positive, the Sign function returns +1, if the
// number is negative, the function returns -1 and if the number is 0 (zero), the function returns 0.
func Sign(number float64) float64 {
	switch {
	case number < 0:
		return -1
	case number == 0:
		return 0
	}
	return 1
}

// Power function returns the result of a given number raised to a supplied power
func Power(number, power float64) (float64, error) {
	if math.IsNaN(number) || math.IsNaN(power) {
		return 0.0, errors.New(errorCode(3))
	}

	return math.Pow(number, power), nil
}

// Product returns the product of a supplied list of numbers
func Product(numbers ...float64) (float64, error) {

	product := 1.0

	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New(errorCode(3))
		}
		product = product * number
	}
	return product, nil
}

// Quotient returns the integer portion of a division between two supplied numbers
func Quotient(numerator, denominator float64) (int, error) {
	if denominator == 0 {
		return 0.0, errors.New(errorCode(2))
	}

	if math.IsNaN(numerator) || math.IsNaN(denominator) {
		return 0.0, errors.New(errorCode(3))
	}

	return int(numerator / denominator), nil
}

// Round rounds a number up or down, to a given number of digits
// The syntax of the function is Round(number, numDigits)
func Round(number float64, numDigits int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(numDigits))
	intermed := number * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

// Sqrt returns the positive square root of a given number
func Sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0.0, errors.New(errorCode(6))
	}

	if math.IsNaN(number) {
		return 0.0, errors.New(errorCode(3))
	}
	return math.Sqrt(number), nil
}

// Sum returns the sum of a supplied list of numbers
func Sum(numbers ...float64) (float64, error) {
	var sum float64
	for _, number := range numbers {
		if math.IsNaN(number) {
			return 0.0, errors.New(errorCode(3))
		}
		sum = sum + number
	}

	return sum, nil
}

// PI Returns the constant value of pi
func PI() float64 {
	return 3.14159265358979
}

// SqrtPI Returns the square root of a supplied number multiplied by pi
// Returns the square root of (number * pi).
func SqrtPI(number float64) (float64, error) {

	// Validate Number - Common Errors
	if number < 0 {
		return 0.0, errors.New(errorCode(6))
	}

	if math.IsNaN(number) {
		return 0.0, errors.New(errorCode(3))
	}

	return math.Sqrt(number * PI()), nil
}

func abs(number float64) float64 {
	switch {
	case number < 0:
		return -number
	case number == 0:
		return 0
	}
	return number
}
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
