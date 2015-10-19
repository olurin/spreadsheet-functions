package financial

import (
	"log"

	"github.com/TaperBox/formulas/lib"
	"github.com/TaperBox/formulas/math"
	"gopkg.in/go-playground/validator.v8"
)

// Basic Description
// =================
//
// The EFFECT function returns the effective annual interest rate
// for a given nominal interest rate and number of compounding periods per year
//
// The format of the function is : Effect( nominal_rate, npery )
//
// where the nominal_rate argument is a numerical value, between 0 and 1, which
// represents the nominal interest rate, and the npery argument is a positive
// integer that gives the number of compounding periods per year
//

// EffectStruct syntax has the following arguments:
// Nominal_rate    Required. The nominal interest rate.
// Npery    Required. The number of compounding periods per year.
type EffectStruct struct {
	Npery       float64 `validate:"required"`
	NominalRate float64 `validate:"gte=0"`
}

// Remarks
// Npery is truncated to an integer.
// If either argument is nonnumeric, EFFECT returns "required" error.
// If nominal_rate ≤ 0 or if npery < 1, EFFECT returns the "Number" error value.
//
func validateEffect(npery, nominalrate float64) (*EffectStruct, error) {

	validate := validator.New(&validator.Config{TagName: "validate"})

	effect := &EffectStruct{
		Npery:       npery,
		NominalRate: nominalrate,
	}

	errs := validate.Struct(effect)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	return effect, nil
}

func (e *EffectStruct) effect() (float64, error) {
	npery := mathlibs.Round(e.Npery, 0)
	p, err := mathlib.Power((1 + (e.NominalRate/100)/npery), npery)
	return p - 1, err
}

// Effect Function The Effective Annual Interest Rate is a measure of
// interest, that incorporates the compounding of interest and is frequently
// used to compare financial loans with different compounding terms.
func Effect(npery, nominalrate float64) (float64, error) {
	value, err := validateEffect(npery, nominalrate)
	if err != nil {
		return 0.0, err
	}
	e, _ := value.effect()
	return e, nil
}
