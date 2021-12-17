// validator_test 用来测试 validator 对嵌套结构体的校验
// 注意：`func (v *Validate) Struct(s interface{}) error` 会自动校验嵌套的结构体
package validator_test

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func printError(err error) {
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
}

type person struct {
	mustHave          // (*validator.Validate).Struct() 函数会自动校验嵌套的结构体
	M        mustHave // 暴露的字段会校验
	m        mustHave // 不暴露的字段不校验
}

type mustHave struct {
	Name string `validate:"required"`
	Age  uint64 `validate:"required"`
}

func Example_validate_struct_1() {
	printError(validate.Struct(&person{}))
	// Output:
	// Key: 'person.mustHave.Name' Error:Field validation for 'Name' failed on the 'required' tag
	// Key: 'person.mustHave.Age' Error:Field validation for 'Age' failed on the 'required' tag
	// Key: 'person.M.Name' Error:Field validation for 'Name' failed on the 'required' tag
	// Key: 'person.M.Age' Error:Field validation for 'Age' failed on the 'required' tag
}

type Foo struct {
	A string
	B int64
}

type bar struct {
	Foo `validate:"structonly,required"`
	F   Foo `validate:"structonly,required"`
}

func Example_validate_struct_2() {
	printError(validate.Struct(&bar{
		Foo: Foo{
			A: "",
			B: 0,
		},
	}))
	// Output:
	// structonly 校验了什么呢？难道是要自定义校验函数？
}

func Example_validate_struct_3() {
	printError(validate.Struct(&bar{}))
	// Output:
	// structonly 校验了什么呢？
}

type req struct {
	F1 struct {
		F1 struct {
			F1 int    `json:"f_1" validate:"required"`
			F2 string `json:"f_2" validate:"required"`
		}
	}
	F2 string `json:"f_2" validate:"required"`
}

func Example_validate_struct_4() {
	printError(validate.Struct(&req{}))
	// Output:
	// Key: 'req.F1.F1.F1' Error:Field validation for 'F1' failed on the 'required' tag
	// Key: 'req.F1.F1.F2' Error:Field validation for 'F2' failed on the 'required' tag
	// Key: 'req.F2' Error:Field validation for 'F2' failed on the 'required' tag
}
