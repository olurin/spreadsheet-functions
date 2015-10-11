package spreadsheet

import (
	"log"

	"gopkg.in/go-playground/validator.v8"
)

// SumStruct struct
type SumStruct struct {
	Numbers []*NumStruct `validate:"required"`
}

// NumStruct struct
type NumStruct struct {
	Number float64 `validate:"required"`
}

func validateSum(numbers []float64) (*SumStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})

	sum := &SumStruct{}

	for _, n := range numbers {
		num := &NumStruct{
			Number: n,
		}
		errs := validate.Struct(num)
		if errs != nil {
			log.Println(num.Number)
			log.Println(errs)
			return nil, errs
		}
		sum.Numbers = append(sum.Numbers, num)
	}

	errs := validate.Struct(sum)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}

	return sum, nil
}

func (nums *SumStruct) sum() float64 {
	var sum float64
	for _, n := range nums.Numbers {
		sum = sum + n.Number
	}
	return sum
}

// Sum function
func Sum(n ...float64) float64 {
	v, err := validateSum(n)
	if err != nil {
		return 0.0
	}

	return v.sum()

}
