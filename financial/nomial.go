package financial

import (
	"log"
	"math"

	"github.com/TaperBox/formulas/math"
	"gopkg.in/go-playground/validator.v8"
)

// Basic Description
// =================
//
// The NOMINAL function returns the nominal annual interest rate,
// given the effective rate and the number of compounding periods per year.
//
// The format of the function is : NOMINAL(effect_rate, npery)

// NominalStruct syntax has the following arguments:
// Effect_rate    Required. The effective interest rate.
// Npery    Required. The number of compounding periods per year.
type NominalStruct struct {
	Npery      float64 `validate:"required"`
	EffectRate float64 `validate:"gte=0"`
}

// Remarks
// Npery is truncated to an integer.
// If either argument is nonnumeric, EFFECT returns "required" error.
// If nominal_rate ≤ 0 or if npery < 1, EFFECT returns the "Number" error value.
//
func validateNominal(npery, effectrate float64) (*NominalStruct, error) {

	validate := validator.New(&validator.Config{TagName: "validate"})

	nominal := &NominalStruct{
		Npery:      npery,
		EffectRate: effectrate,
	}

	errs := validate.Struct(nominal)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	return nominal, nil
}

func (e *NominalStruct) nominal() float64 {
	npery := mathlib.Round(e.Npery, 0)
	return npery * (math.Pow(((e.EffectRate/100)+1), 1.0/npery) - 1)

}

// Nominal returns the nominal interest rate for a given effective
// interest rate and number of compounding periods per year.
func Nominal(npery, nominalrate float64) float64 {
	value, err := validateNominal(npery, nominalrate)
	if err != nil {
		return 0.0
	}
	return value.nominal()
}
