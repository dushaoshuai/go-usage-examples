package validator_test

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type person struct {
	mustHave
	Money int64
}

type mustHave struct {
	Name string `validate:"required"`
	Age  uint64 `validate:"required"`
}

func ExampleValidateStruct() {
	person := person{
		mustHave: mustHave{
			Name: "",
			Age:  0,
		},
		Money: 0,
	}
	err := validate.Struct(person)
	// https://raw.githubusercontent.com/go-playground/validator/master/_examples/struct-level/main.go
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		for _, err = range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	}
	// Output:
	// Key: 'person.mustHave.Name' Error:Field validation for 'Name' failed on the 'required' tag
	// Key: 'person.mustHave.Age' Error:Field validation for 'Age' failed on the 'required' tag
}
