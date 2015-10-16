package spreadsheet

import (
	"log"

	"gopkg.in/go-playground/validator.v8"
)

// QuotientStruct struct
type QuotientStruct struct {
	Numerator   float64 `validate:"required"`
	Denominator float64 `validate:"required"`
}

func validateQuotient(num, den float64) (*QuotientStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})

	quo := &QuotientStruct{
		Numerator:   num,
		Denominator: den,
	}

	errs := validate.Struct(quo)
	if errs != nil {
		return nil, errs
	}
	return quo, nil
}

// Quotient function
func Quotient(numerator, denominator float64) int {
	v, err := validateQuotient(numerator, denominator)
	if err != nil {
		log.Println(err)
		return 0
	}

	return int(v.Numerator / v.Denominator)
}
