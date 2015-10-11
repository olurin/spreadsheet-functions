package spreadsheet

import (
	"log"
	"math"

	"gopkg.in/go-playground/validator.v8"
)

//Log10Struct struct
type Log10Struct struct {
	Number float64 `validate:"required"`
}

func validateLog10(num float64) (*Log10Struct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	ln := &Log10Struct{
		Number: num,
	}

	errs := validate.Struct(ln)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}
	return ln, nil
}

// Log10 Returns the base-10 logarithm of a number.
func Log10(num float64) float64 {
	v, err := validateLog10(num)
	if err != nil {
		return 0.0
	}

	return math.Log10(v.Number)
}
