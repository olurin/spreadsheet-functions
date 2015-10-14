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

// SqrtPI Returns the square root of a supplied number multiplied by pi
// Returns the square root of (number * pi).
func SqrtPI(number float64) float64 {
	v, err := validateTrigonometry(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return spreadsheet.Sqrt(v.Number * PI())
}

// Degree Converts radians into degrees.
func Degree(radians float64) float64 {
	v, err := validateTrigonometry(radians)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return v.Number * (180 / PI())
}

// Radians Converts Degrees to Radians
func Radians(degree float64) float64 {
	v, err := validateTrigonometry(degree)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return v.Number * (PI() / 180)
}
