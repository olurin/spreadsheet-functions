package spreadsheet

import (
	"errors"
	"log"

	"gopkg.in/go-playground/validator.v8"
)

// The GCD function returns the greatest common divisor of two or more supplied integers.
// The format of the function is :   GCD( number1, number2, ...)
// Where the number arguments are up to 255 numerical values for which you want to calculate the greatest common divisor.
// If any of the supplied numbers are not integers, these values are truncated to integers.

// GCDstruct struct
type GCDstruct struct {
	Numbers []int `validate:"required"`
}

func validateGCD(nums []int) (*GCDstruct, error) {

	validate := validator.New(&validator.Config{TagName: "validate"})

	if len(nums) <= 0 {
		return nil, errors.New("Number1, number2, ...    Number1 is required, subsequent numbers are optional. 1 to 255 values. If any value is not an integer, it is truncated.")
	}

	gcd := &GCDstruct{}

	for _, number := range nums {
		gcd.Numbers = append(gcd.Numbers, number)
	}

	errs := validate.Struct(gcd)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	return gcd, nil
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

// GCD test function
func GCD(numbers ...int) int {

	value, err := validateGCD(numbers)
	if err != nil {
		return 0
	}

	var gcdValue int

	l := len(value.Numbers) - 1
	for l >= 0 {
		if gcdValue == 0 {
			gcdValue = gcd(value.Numbers[l-1], value.Numbers[l])
			l = l - 1
			continue
		}
		gcdValue = gcd(value.Numbers[l], gcdValue)
		l = l - 1
	}

	return gcdValue
}
