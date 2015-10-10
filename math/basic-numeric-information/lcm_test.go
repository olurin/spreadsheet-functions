package spreadsheet

import (
	"errors"
	"log"

	"gopkg.in/go-playground/validator.v8"
)

// LCMstruct struct
type LCMstruct struct {
	Numbers []int `validate:"required"`
}

func validateLCM(nums []int) (*LCMstruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})

	if len(nums) <= 0 {
		return nil, errors.New("Number1, number2, ...    Number1 is required, subsequent numbers are optional. 1 to 255 values. If any value is not an integer, it is truncated.")
	}

	lcm := &LCMstruct{}

	for _, number := range nums {
		lcm.Numbers = append(lcm.Numbers, number)
	}

	errs := validate.Struct(lcm)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	return lcm, nil
}
