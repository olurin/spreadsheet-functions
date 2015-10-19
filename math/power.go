package mathlib

import (
	"log"
	"math"

	"gopkg.in/go-playground/validator.v8"
)

// PowerStruct struct
type PowerStruct struct {
	Number float64 `validate:"required"`
	Power  float64 `validate:"required"`
}

func validatePower(number, power float64) (*PowerStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})
	pw := &PowerStruct{
		Number: number,
		Power:  power,
	}

	errs := validate.Struct(pw)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}
	return pw, nil
}

func (power *PowerStruct) power() float64 {
	return math.Pow(power.Number, power.Power)
}

// Power function returns the result of a given number raised to a supplied power
func Power(number, power float64) (float64, error) {
	v, err := validatePower(number, power)
	if err != nil {
		return 0.0, err
	}
	return v.power(), nil
}
