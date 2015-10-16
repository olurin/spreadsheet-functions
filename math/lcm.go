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

// LCM function
func LCM(numbers ...int) int {

	value, err := validateLCM(numbers)
	if err != nil {
		return 0
	}

	var lcmvalue int
	var index = 0

	l := len(value.Numbers) - 1

	for l >= 0 {

		if index == 0 {
			lcmvalue = (value.Numbers[l-1] * value.Numbers[l]) / gcd(value.Numbers[l-1], value.Numbers[l])
			l = l - 1
			index = index + 1
			continue
		}

		lcmvalue = (value.Numbers[l] * lcmvalue) / gcd(value.Numbers[l], lcmvalue)
		l = l - 1
	}
	return lcmvalue
}
