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
func Degree(radian float64) (float64, error) { return mathlib.Degree(radians) }

// Radians converts Degrees to Radians
func Radians(degree float64) (float64, error) { return mathlib.Radians(degree) }

// Cos returns the Cosine of a given angle
func Cos(degree float64) (float64, error) { return mathlib.Cos(degree) }

//Acos returns the Arccosine of a number
func Acos(number float64) (float64, error) { return mathlib.Acos(degree) }
