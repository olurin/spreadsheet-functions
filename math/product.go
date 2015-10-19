package mathlib

import (
	"log"

	"gopkg.in/go-playground/validator.v8"
)

// ProductStruct struct
type ProductStruct struct {
	Numbers []*NumStruct `validate:"required"`
}

func validateProduct(numbers []float64) (*ProductStruct, error) {
	validate := validator.New(&validator.Config{TagName: "validate"})

	product := &ProductStruct{}

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
		product.Numbers = append(product.Numbers, num)
	}

	errs := validate.Struct(product)
	if errs != nil {
		log.Println(errs)
		return nil, errs
	}
	return product, nil
}

func (nums *ProductStruct) product() float64 {
	prod := 1.0
	for _, n := range nums.Numbers {
		prod = prod * n.Number
	}
	return prod
}

// Product function
func Product(p ...float64) (float64, error) {
	v, err := validateProduct(p)
	if err != nil {
		return 0.0, err
	}
	return v.product(), nil
}
