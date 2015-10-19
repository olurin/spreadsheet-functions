package mathlib

import (
	"log"
	"math"

	"gopkg.in/go-playground/validator.v8"
)

// TrigonometryStruct struct
type TrigonometryStruct struct {
	Number float64
}

func validateTrigonometry(number float64) (*TrigonometryStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	tri := &TrigonometryStruct{
		Number: number,
	}

	errs := validate.Struct(tri)
	if errs != nil {
		return nil, errs
	}

	return tri, nil
}

// PI Returns the constant value of pi
func PI() float64 {
	return 3.14159265358979
}

// SqrtPI Returns the square root of a supplied number multiplied by pi
// Returns the square root of (number * pi).
func SqrtPI(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return Sqrt(v.Number * PI())
}

// Degree Converts radians into degrees.
func Degree(radians float64) float64 {
	v, err := validateTrigonometry(radians)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return v.Number * (180 / PI())
}

// Radians Converts Degrees to Radians
func Radians(degree float64) float64 {
	v, err := validateTrigonometry(degree)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return v.Number * (PI() / 180)
}

// Cos returns the Cosine of a given angle
func Cos(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	// Math returns the cosine of the radian argument x.
	return math.Cos(v.Number)
}

//Acos returns the Arccosine of a number
func Acos(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}

	return math.Acos(v.Number)
}

//CosH returns the hyperbolic cosine of x
func CosH(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}

	return math.Cosh(v.Number)
}

//AcosH returns the inverse hyperbolic cosine of x
func AcosH(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}

	return math.Acosh(v.Number)
}

//Sec returns the secant of an angle (New in Excel 2013)
func Sec(number float64) float64 {
	return (1 / Cos(number))
}

//AsecH returns the hyperbolic secant of an angle
func AsecH(number float64) float64 {
	return (1 / CosH(number))
}

// Sin returns the Sine of a given angle
func Sin(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	// Math returns the sine of the radian argument x.
	return math.Sin(v.Number)
}

// Asin returns the arcsine, in radians, of x
func Asin(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}

	return math.Asin(v.Number)
}

// SinH returns the Cosine of a given angle
func SinH(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	// Sinh returns the hyperbolic sine of x.
	return math.Sinh(v.Number)
}

// AsinH function returns the inverse hyperbolic sine. Therefore: ASINH( SINH( z ) ) = z
func AsinH(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return math.Asinh(v.Number)
}

// Csc returns the cosecant of an angle
func Csc(number float64) float64 {
	return 1 / Sin(number)
}

//CscH Returns the hyperbolic cosecant of an angle
func CscH(number float64) float64 {
	return 1 / SinH(number)
}

// Tan function calculates the tangent of a given angle
func Tan(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return math.Tan(v.Number)
}

// Atan returns the Arctangent of a given number
func Atan(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return math.Atan(v.Number)
}

//Atan2 Returns the Arctangent of a given pair of x and y coordinates
func Atan2(x, y float64) float64 {
	xValue, err := validateTrigonometry(x)
	if err != nil {
		log.Println(err)
		return 0.0
	}

	yValue, err := validateTrigonometry(y)
	if err != nil {
		log.Println(err)
		return 0.0
	}

	return math.Atan2(xValue.Number, yValue.Number)
}

//TanH Returns the Hyperbolic Tangent of a given number
func TanH(number float64) float64 {
	return SinH(number) / CosH(number)
}

// AtanH function calculates the inverse hyperbolic tangent of a supplied number
func AtanH(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return math.Atanh(v.Number)
}

// Cot returns the cotangent of an angle
func Cot(number float64) float64 {
	return 1 / Tan(number)
}

//CotH returns the hyperbolic cotangent of an angle
func CotH(number float64) float64 {
	return 1 / TanH(number)
}
