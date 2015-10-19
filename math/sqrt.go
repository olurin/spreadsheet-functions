package mathlib

import (
	"log"
	"math"

	"gopkg.in/go-playground/validator.v8"
)

// SqrtStruct struct
type SqrtStruct struct {
	Number float64 `validate:"required"`
}

func validateSqrt(number float64) (*SqrtStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	sq := &SqrtStruct{
		Number: number,
	}
	errs := validate.Struct(sq)
	if errs != nil {
		return nil, errs
	}
	return sq, nil
}

// Sqrt function
func Sqrt(number float64) float64 {
	v, err := validateSqrt(number)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return math.Sqrt(v.Number)
}
