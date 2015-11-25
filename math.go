package formulas

import (
	"log"
	"math"
)

// ABS function returns the absolute value of a number
// The absolute value of a number is the number without its sign
func ABS(number float64) (float64, FormulaErrors) {
	// Validate Number
	if math.IsNaN(number) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "The supplied number argument cannot be recognised as numeric value",
			Type:    ErrorExport(3),
			Formula: "ABS",
		}
		return 0.0, err
	}

	return abs(number), nil
}

// Exp function returns e raised to a given power
func Exp(number float64) (float64, FormulaErrors) {

	// Validate Number
	if math.IsNaN(number) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "The supplied number argument cannot be recognised as numeric value",
			Type:    ErrorExport(3),
			Formula: "Exp",
		}
	}

	return math.Exp(number), nil
}

// GCD function returns the greatest common divisor of two or more supplied integers.
// The format of the function is :   GCD( number1, number2, ...)
// Where the number arguments are up to 255 numerical values for which you want to calculate the greatest common divisor.
// If any of the supplied numbers are not integers, these values are truncated to integers.
func GCD(numbers ...int) (int, FormulaErrors) {

	// Validate Numbers
	if len(numbers) <= 1 {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "Number1, number2, ...    Number1 is required, subsequent numbers are optional. 1 to 255 values. If any value is not an integer, it is truncated.",
			Type:    ErrorExport(6),
			Formula: "GCD",
		}

		return 0, err
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
func LCM(numbers ...int) (int, FormulaErrors) {

	// Validate Numbers
	if len(numbers) <= 1 {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "Number1, number2, ...    Number1 is required, subsequent numbers are optional. 1 to 255 values. If any value is not an integer, it is truncated.",
			Type:    ErrorExport(6),
			Formula: "LCM",
		}
		return 0, err
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
func LN(number float64) (float64, FormulaErrors) {
	// Validate Number
	if number <= 0 {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied number argument is negative or zero",
			Type:    ErrorExport(6),
			Formula: "LN",
		}
		return 0.0, err
	}

	if math.IsNaN(number) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied number argument is not a numeric value",
			Type:    ErrorExport(3),
			Formula: "LN",
		}
		return 0.0, err
	}

	return math.Log(number), nil
}

// Log10 Returns the base-10 logarithm of a number.
func Log10(number float64) (float64, FormulaErrors) {

	// Validate Number
	if number <= 0 {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied number argument is negative or zero",
			Type:    ErrorExport(6),
			Formula: "Log10",
		}
		return 0.0, err
	}

	if math.IsNaN(number) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied number argument is not a numeric value",
			Type:    ErrorExport(3),
			Formula: "Log10",
		}
		return 0.0, err
	}

	return math.Log10(number), nil
}

// Mod returns the remainder after number is divided by divisor. The result has the same sign as divisor.
func Mod(number, divisor float64) (float64, FormulaErrors) {

	if math.IsNaN(divisor) || math.IsNaN(number) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied arguments are non-numeric",
			Type:    ErrorExport(6),
			Formula: "Mod",
		}
		return 0.0, err
	}

	if divisor == 0 {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied denominator argument is zero",
			Type:    ErrorExport(2), // #DIV/0!
			Formula: "Mod",
		}
		return 0.0, err
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
func Power(number, power float64) (float64, FormulaErrors) {
	if math.IsNaN(number) || math.IsNaN(power) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "supplied arguments are non-numeric",
			Type:    ErrorExport(3),
			Formula: "Power",
		}
		return 0.0, err
	}

	return math.Pow(number, power), nil
}

// Product returns the product of a supplied list of numbers
func Product(numbers ...float64) (float64, FormulaErrors) {

	product := 1.0

	for _, number := range numbers {
		if math.IsNaN(number) {
			var err FormulaErrors
			err["Error"] = &FunctionError{
				Error:   "Arguments are supplied directly to the Product function can not be interpreted as numeric values",
				Type:    ErrorExport(3),
				Formula: "Product",
			}
			return 0.0, err
		}
		product = product * number
	}
	return product, nil
}

// Quotient returns the integer portion of a division between two supplied numbers
func Quotient(numerator, denominator float64) (int, FormulaErrors) {
	if denominator == 0 {

		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "The supplied denominator argument is zero",
			Type:    ErrorExport(2),
			Formula: "Quotient",
		}

		return 0.0, err
	}

	if math.IsNaN(numerator) || math.IsNaN(denominator) {

		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "The supplied arguments are non-numeric",
			Type:    ErrorExport(3),
			Formula: "Quotient",
		}
		return 0.0, err
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
func Sqrt(number float64) (float64, FormulaErrors) {
	if number < 0 {
		err := FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "The supplied number argument is negative",
			Type:    ErrorExport(6),
			Formula: "Sqrt",
		}

		log.Println(err)
		return 0.0, err
	}

	if math.IsNaN(number) {
		var err FormulaErrors
		err["Error"] = &FunctionError{
			Error:   "The supplied number argument is non-numeric",
			Type:    ErrorExport(3),
			Formula: "Sqrt",
		}
		return 0.0, err
	}

	return math.Sqrt(number), nil
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
