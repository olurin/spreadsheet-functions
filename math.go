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

// Degrees Converts radians into degrees.
func Degrees(radians float64) (float64, error) {

	if math.IsNaN(radians) {
		return 0.0, errors.New(errorCode(3))
	}

	return radians * (180 / PI()), nil
}

// Radians Converts Degrees to Radians
func Radians(degrees float64) (float64, error) {

	if math.IsNaN(degrees) {
		return 0.0, errors.New(errorCode(3))
	}

	return degrees * (PI() / 180), nil
}

// Cos returns the Cosine of a given angle
func Cos(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	// Math returns the cosine of the radian argument x.
	return math.Cos(number), nil
}

//Acos returns the Arccosine of a number
func Acos(number float64) (float64, error) {

	// Validate Number - Common Errors
	if number < -1 || number > 1 {
		return 0.0, errors.New("#NUM!	-	Occurred because the supplied number argument is outside of the range -1 to +1")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Acos(number), nil
}

//CosH returns the hyperbolic cosine of x
func CosH(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Cosh(number), nil
}

//AcosH returns the inverse hyperbolic cosine of x
func AcosH(number float64) (float64, error) {

	if number < 1 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than 1")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Acosh(number), nil
}

//Sec returns the secant of an angle
func Sec(number float64) (float64, error) {

	if number < -134217728 || number > 134217728 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than -2^27 or is greater than 2^27")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return (1 / math.Cos(number)), nil
}

//SecH returns the hyperbolic secant of an angle
func SecH(number float64) (float64, error) {

	if number < -134217728 || number > 134217728 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than -2^27 or is greater than 2^27")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return (1 / math.Cosh(number)), nil
}

// Sin returns the Sine of a given angle
func Sin(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	// Math returns the sine of the radian argument x.
	return math.Sin(number), nil
}

// Asin returns the arcsine, in radians, of x
func Asin(number float64) (float64, error) {

	// Validate Number - Common Errors
	if number < -1 || number > 1 {
		return 0.0, errors.New("#NUM!	-	Occurred because the supplied number argument is outside of the range -1 to +1")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Asin(number), nil
}

// SinH returns the Cosine of a given angle
func SinH(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	// Sinh returns the hyperbolic sine of x.
	return math.Sinh(number), nil
}

// AsinH function returns the inverse hyperbolic sine. Therefore: ASINH( SINH( z ) ) = z
func AsinH(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Asinh(number), nil
}

// Csc returns the cosecant of an angle
func Csc(number float64) (float64, error) {

	if number == 0 {
		return 0.0, errors.New("#DIV/0!	-	Occurred because the supplied number argument is equal to zero")
	}

	if number < -134217728 || number > 134217728 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than -2^27 or is greater than 2^27")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return 1 / math.Sin(number), nil
}

//CscH Returns the hyperbolic cosecant of an angle
func CscH(number float64) (float64, error) {

	if number == 0 {
		return 0.0, errors.New("#DIV/0!	-	Occurred because the supplied number argument is equal to zero")
	}

	if number < -134217728 || number > 134217728 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than -2^27 or is greater than 2^27")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return 1 / math.Sinh(number), nil
}

// Tan function calculates the tangent of a given angle
func Tan(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Tan(number), nil
}

// Atan returns the Arctangent of a given number
func Atan(number float64) (float64, error) {
	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}
	return math.Atan(number), nil
}

//Atan2 Returns the Arctangent of a given pair of x and y coordinates
func Atan2(x, y float64) (float64, error) {

	if math.IsNaN(x) || math.IsNaN(y) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	if x == 0 && y == 0 {
		return 0.0, errors.New("#DIV/0!	-	Occurred because the supplied x_num and y_num arguments are both equal to 0")
	}

	return math.Atan2(x, y), nil
}

//TanH Returns the Hyperbolic Tangent of a given number
func TanH(number float64) (float64, error) {
	sinh, err := SinH(number)
	if err != nil {
		return 0.0, err
	}

	cosh, err := CosH(number)
	if err != nil {
		return 0.0, err
	}

	return sinh / cosh, nil
}

// AtanH function calculates the inverse hyperbolic tangent of a supplied number
func AtanH(number float64) (float64, error) {

	// Validate Number - Common Errors
	if number <= -1 || number >= 1 {
		return 0.0, errors.New("#NUM!	-	Occurred because the supplied number argument is  ≤ -1 or ≥ 1")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Atanh(number), nil
}

// Cot returns the cotangent of an angle
func Cot(number float64) (float64, error) {

	if number == 0 {
		return 0.0, errors.New("#DIV/0!	-	Occurred because the supplied number argument is equal to zero")
	}

	if number < -134217728 || number > 134217728 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than -2^27 or is greater than 2^27")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return 1 / math.Tan(number), nil
}

//CotH returns the hyperbolic cotangent of an angle
func CotH(number float64) (float64, error) {

	if number == 0 {
		return 0.0, errors.New("#DIV/0!	-	Occurred because the supplied number argument is equal to zero")
	}

	if number < -134217728 || number > 134217728 {
		return 0.0, errors.New("#NUM! -  Occurred because the supplied number argument is less than -2^27 or is greater than 2^27")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	tanh, err := TanH(number)
	if err != nil {
		return 0.0, err
	}

	return (1 / tanh), nil
}

// Acot (Arccotangent) returns the arccotangent of a number
// Inverse cotangent in radians
func Acot(number float64) (float64, error) {

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	number = 1 / number

	v, err := Atan(number)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

//AcotH returns the hyperbolic arccotangent of a number
func AcotH(number float64) (float64, error) {

	// Validate Number - Common Errors
	if number >= -1 && number <= 1 {
		return 0.0, errors.New("#NUM!	-	Occurred because supplied number argument is between -1 and +1 (inclusive)")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	v, err := LN((number + 1) / (number - 1))
	if err != nil {
		return 0.0, err
	}
	return (0.5 * v), nil
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
