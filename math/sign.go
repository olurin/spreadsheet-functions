package mathlib

// Sign determines the sign of a number. Returns 1 if the number is positive,
// zero (0) if the number is 0, and -1 if the number is negative.

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

// SignStruct syntax has the following arguments:
// Number    Required. Any real number.
type SignStruct struct {
	Number float64 `validate:"required"`
}

func validateSign(number float64) (*SignStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})

	sign := &SignStruct{
		Number: number,
	}

	errs := validate.Struct(sign)
	if errs != nil {
		fmt.Println(errs)
		return nil, errs
	}

	return sign, nil
}

func (number *SignStruct) sign() float64 {
	switch {
	case number.Number < 0:
		return -1
	case number.Number == 0:
		return 0
	}
	return 1
}

// Sign function returns the sign (+1, -1 or 0) of a supplied number. ie.
// if the number is positive, the Sign function returns +1, if the
// number is negative, the function returns -1 and if the number is 0 (zero), the function returns 0.
func Sign(number float64) float64 {
	value, err := validateSign(number)
	if err != nil {
		return number
	}
	return value.sign()
}
