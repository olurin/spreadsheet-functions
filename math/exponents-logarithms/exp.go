package spreadsheet

import (
	"log"
	"math"

	"gopkg.in/go-playground/validator.v8"
)

// The Excel EXP function calculates the value of the mathematical constant e, raised to the power of a given number.
// The format of the function is:
// EXP( number )

// ExpStruct struct
type ExpStruct struct {
	Number float64 `validate:"required"`
}

func validateEXP(num float64) (*ExpStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	exp := &ExpStruct{
		Number: num,
	}

	errs := validate.Struct(exp)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}
	return exp, nil
}

// Exp function
func Exp(num float64) float64 {
	v, err := validateEXP(num)
	if err != nil {
		return 0.0
	}

	return math.Exp(v.Number)
}
