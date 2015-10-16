package spreadsheet

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

// The ABS function returns the absolute value
// Pretty much a validated wrapper for Math.Abs
// (ie. the modulus) of any supplied number
//
// The Syntax of the function is: ABS(number)
//  ABS(number)

// AbsStruct function syntax has the following arguments:
// Number    Required. The real number of which you want the absolute value.
type AbsStruct struct {
	Number float64 `validate:"required"`
}

// Validate ABS function
// If you get an error from the function, this is going to return a nil struct and error
func validateAbs(number float64) (*AbsStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})

	abs := &AbsStruct{
		Number: number,
	}

	errs := validate.Struct(abs)
	if errs != nil {
		fmt.Println(errs)
		return nil, errs
	}

	return abs, nil
}

func (number *AbsStruct) abs() float64 {
	switch {
	case number.Number < 0:
		return -number.Number
	case number.Number == 0:
		return 0
	}
	return number.Number
}

// Abs function returns the absolute value of a number.
// The absolute value of a number is the number without its sign.
func Abs(number float64) float64 {
	value, err := validateAbs(number)
	if err != nil {
		return number
	}
	return value.abs()
}
