package mathlib

import (
	"errors"
	"math"
)

// PI Returns the constant value of pi
func PI() float64 {
	return 3.14159265358979
}

// SqrtPI Returns the square root of a supplied number multiplied by pi
// Returns the square root of (number * pi).
func SqrtPI(number float64) (float64, error) {

	// Validate Number - Common Errors
	if number < 0 {
		return 0.0, errors.New("#NUM!	-	Occurred because the supplied number argument is < 0")
	}

	if math.IsNaN(number) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return math.Sqrt(number * PI()), nil
}

// Degree Converts radians into degrees.
func Degree(radians float64) (float64, error) {

	if math.IsNaN(radians) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return radians * (180 / PI()), nil
}

// Radians Converts Degrees to Radians
func Radians(degree float64) (float64, error) {

	if math.IsNaN(degree) {
		return 0.0, errors.New("#VALUE!	-	Occurred because the supplied number argument is non-numeric")
	}

	return degree * (PI() / 180), nil
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
