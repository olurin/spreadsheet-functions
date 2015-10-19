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
