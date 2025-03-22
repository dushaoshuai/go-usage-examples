// validator_test 用来测试 validator 对嵌套结构体的校验
// 注意：`func (v *Validate) Struct(s interface{}) error` 会自动校验嵌套的结构体
package validator_test

import (
	"fmt"
	"slices"
	"testing"
	"time"

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
	Name     string    `validate:"required"`
	Age      uint64    `validate:"required"`
	Birthday time.Time `validate:"required"` // 时间也可以校验
}

func Example_validate_struct_1() {
	printError(validate.Struct(&person{}))
	// Output:
	// Key: 'person.mustHave.Name' Error:Field validation for 'Name' failed on the 'required' tag
	// Key: 'person.mustHave.Age' Error:Field validation for 'Age' failed on the 'required' tag
	// Key: 'person.mustHave.Birthday' Error:Field validation for 'Birthday' failed on the 'required' tag
	// Key: 'person.M.Name' Error:Field validation for 'Name' failed on the 'required' tag
	// Key: 'person.M.Age' Error:Field validation for 'Age' failed on the 'required' tag
	// Key: 'person.M.Birthday' Error:Field validation for 'Birthday' failed on the 'required' tag
}

type Foo struct {
	A string
	B int64
}

type bar struct {
	Foo `validate:"required"`
	F   Foo `validate:"required"`
}

func TestRequiredStruct(t *testing.T) {
	v := validator.New(validator.WithRequiredStructEnabled())

	tests := []struct {
		b       bar
		wantErr bool
	}{
		{
			b:       bar{},
			wantErr: true,
		},
		{
			b: bar{
				Foo: Foo{
					A: "a",
				},
			},
			wantErr: true,
		},
		{
			b: bar{
				F: Foo{
					A: "a",
					B: 0,
				},
			},
			wantErr: true,
		},
		{
			b: bar{
				Foo: Foo{
					A: "a",
				},
				F: Foo{
					A: "a",
					B: 0,
				},
			},
			wantErr: false,
		},
	}

	for test := range slices.Values(tests) {
		err := v.Struct(test.b)
		if (err != nil) != test.wantErr {
			t.Errorf("validateStruct() error = %v, wantErr %v", err, test.wantErr)
		}
	}
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

type cat struct {
	Name string `json:"name" validate:"required"`
	Age  int64  `json:"age" validate:"gte=0"`
}

type cats struct {
	A uint64
	C []cat `json:"c" validate:"required_if=A 1,dive"` // 如果 C 是空切片，不会 dive
}

func Example_dive_into_slice() {
	cs1 := cats{
		A: 0,
		C: []cat{
			{"二狗", -56},
			{"", 0},
		}}
	cs2 := cats{
		A: 0,
		C: nil,
	}
	printError(validate.Struct(cs1))
	printError(validate.Struct(cs2))

	// Output:
	// Key: 'cats.C[0].Age' Error:Field validation for 'Age' failed on the 'gte' tag
	// Key: 'cats.C[1].Name' Error:Field validation for 'Name' failed on the 'required' tag
}

type required_if_And_oneof struct {
	A uint64
	B string `validate:"required_if=A 10,oneof=Golang Java C Lua"` // required_if 不成立，oneof 仍然会检查
}

func Example_required_if_And_oneof() {
	printError(validate.Struct(required_if_And_oneof{
		A: 0,
		B: "Perl",
	}))
	// Output:
	// Key: 'required_if_And_oneof.B' Error:Field validation for 'B' failed on the 'oneof' tag
}

func Example_required() {
	type coo struct {
		C int
	}
	type foo struct {
		A []coo `validate:"required"` // 对切片进行校验，还是 `validate:"required,gte=1"` 比较好，
	} // 因为 required 校验的是切片不是 nil，切片的长度仍可能是 0
	printError(validate.Struct(foo{
		A: []coo{}, //
	}))
	// Output:
}
