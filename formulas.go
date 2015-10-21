package formulas

import "github.com/TaperBox/formulas/math"

// Abs returns the absolute value (ie. the modulus) of a supplied number
func Abs(number float64) (float64, error) { return mathlib.Abs(number) }

// Exp returns e raised to a given power
func Exp(number float64) (float64, error) { return mathlib.Exp(number) }

// GCD returns the Greatest Common Divisor of two or more supplied numbers
func GCD(numbers []int) (int, error) { return mathlib.GCD(numbers) }

// LCM returns the Least Common Multiple of two or more supplied numbers
func LCM(numbers []int) (int, error) { return mathlib.LCM(numbers) }

// LN returns the natural logarithm of a given number
func LN(number float64) (float64, error) { return mathlib.LN(number) }

// Log10 returns the base 10 logarithm of a given number
func Log10(number float64) (float64, error) { return mathlib.Log10(number) }

// Mod returns the remainder after number is divided by divisor. The result has the same sign as divisor
func Mod(number, divisor float64) (float64, error) { return mathlib.Mod(number, divisor) }

// Power returns the result of a number raised to a power
func Power(number, power float64) (float64, error) { return mathlib.Power(number, power) }

// Round rounds a number to a specified number of digits
func Round(number float64, numDigits int) float64 { return mathlib.Round(number, numDigits) }

// PI Returns the constant value of pi
func PI() float64 { return 3.14159265358979 }

// SqrtPI returns the square root of a supplied number multiplied by pi
func SqrtPI(number float64) (float64, error) { return mathlib.SqrtPI(number) }

// Degree converts radians into degrees.
func Degree(radian float64) (float64, error) { return mathlib.Degree(radian) }

// Radians converts Degrees to Radians
func Radians(degree float64) (float64, error) { return mathlib.Radians(degree) }

// Cos returns the Cosine of a given angle
func Cos(degree float64) (float64, error) { return mathlib.Cos(degree) }

//Acos returns the Arccosine of a number
func Acos(number float64) (float64, error) { return mathlib.Acos(number) }

//CosH returns the hyperbolic cosine of x
func CosH(number float64) (float64, error) { return mathlib.CosH(number) }

//AcosH returns the inverse hyperbolic cosine of x
func AcosH(number float64) (float64, error) { return mathlib.AcosH(number) }

//Sec returns the secant of an angle
func Sec(number float64) (float64, error) { return mathlib.Sec(number) }

//SecH returns the hyperbolic secant of an angle
func SecH(number float64) (float64, error) { return mathlib.SecH(number) }

// Sin returns the Sine of a given angle
func Sin(number float64) (float64, error) { return mathlib.Sin(number) }

// Asin returns the arcsine, in radians, of x
func Asin(number float64) (float64, error) { return mathlib.Asin(number) }

// SinH returns the Cosine of a given angle
func SinH(number float64) (float64, error) { return mathlib.SinH(number) }

// AsinH function returns the inverse hyperbolic sine
func AsinH(number float64) (float64, error) { return mathlib.AsinH(number) }

// Csc returns the cosecant of an angle
func Csc(number float64) (float64, error) { return mathlib.Csc(number) }

//CscH Returns the hyperbolic cosecant of an angle
func CscH(number float64) (float64, error) { return mathlib.CscH(number) }

// Tan function calculates the tangent of a given angle
func Tan(number float64) (float64, error) { return mathlib.TanH(number) }

// Atan returns the Arctangent of a given number
func Atan(number float64) (float64, error) { return mathlib.Atan(number) }

//Atan2 returns the Arctangent of a given pair of x and y coordinates
func Atan2(x, y float64) (float64, error) { return mathlib.Atan2(x, y) }

//TanH Returns the Hyperbolic Tangent of a given number
func TanH(number float64) (float64, error) { return mathlib.TanH(number) }

// AtanH function calculates the inverse hyperbolic tangent of a supplied number
func AtanH(number float64) (float64, error) { return mathlib.AtanH(number) }

// Cot returns the cotangent of an angle
func Cot(number float64) (float64, error) { return mathlib.Cot(number) }

// Acot (Arccotangent) returns the arccotangent of a number
// Inverse cotangent in radians
func Acot(number float64) (float64, error) { return mathlib.AcotH(number) }

//AcotH returns the hyperbolic arccotangent of a number
func AcotH(number float64) (float64, error) { return mathlib.AcotH(number) }
