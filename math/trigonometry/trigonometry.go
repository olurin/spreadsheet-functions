package spreadsheet

import (
	"log"

	"github.com/TaperBox/spreadsheet-functions/math/basic-mathematical-operations"
	"gopkg.in/go-playground/validator.v8"
)

// TrigonometryStruct struct
type TrigonometryStruct struct {
	Number float64 `validate:"required"`
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

// SQRTPI Returns the square root of a supplied number multiplied by pi
// Returns the square root of (number * pi).
func SQRTPI(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return spreadsheet.Sqrt(v.Number * PI())
}
