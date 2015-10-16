package spreadsheet

// LN(number)
// The LN function syntax has the following arguments:
// Number    Required. The positive real number for which you want the natural logarithm.
// LN is the inverse of the EXP function.

import (
	"log"
	"math"

	"gopkg.in/go-playground/validator.v8"
)

// LNStruct struct
type LNStruct struct {
	Number float64 `validate:"required"`
}

func validateLN(num float64) (*LNStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	ln := &LNStruct{
		Number: num,
	}

	errs := validate.Struct(ln)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}
	return ln, nil
}

// LN returns the natural logarithm of a number. Natural logarithms are based on the constant e (2.71828182845904)
func LN(num float64) float64 {
	v, err := validateLN(num)
	if err != nil {
		return 0.0
	}

	return math.Log(v.Number)
}
