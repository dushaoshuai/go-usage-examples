package validate

import (
	"net/http"
	"testing"

	validator "github.com/go-playground/validator/v10"
	"github.com/samber/lo"
)

// User contains user information
type User struct {
	FirstName      string     `json:"first_name" validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `json:"age" zh_Hans_CN:"年龄" validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `json:"favourite_color" validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"`         // a person can have a home and cottage...
	Books          []*Book    `validate:"min=10,dive,required"`
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

type Book struct {
	Name string `json:"name" validate:"required"`
}

func TestValidate(t *testing.T) {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	err := Validator{}.Validate((*http.Request)(nil), user)
	if err != nil {
		t.Fatal(err)
	}
}

type requiredWithout struct {
	ThresholdUpperLimit *float64 `json:"threshold_upper_limit" zh_Hans_CN:"人工阈值上限" validate:"required_without=ThresholdLowerLimit"` // 人工阈值上限
	ThresholdLowerLimit *float64 `json:"threshold_lower_limit" zh_Hans_CN:"人工阈值下限" validate:"required_without=ThresholdUpperLimit"` // 人工阈值下限
}

func Test_required_without(t *testing.T) {
	tests := []struct {
		name    string
		args    requiredWithout
		wantErr bool
	}{
		{
			name:    "both_nil",
			args:    requiredWithout{},
			wantErr: true,
		},
		{
			name: "upper_nil",
			args: requiredWithout{
				ThresholdUpperLimit: nil,
				ThresholdLowerLimit: lo.ToPtr(6.10),
			},
			wantErr: false,
		},
		{
			name: "lower_nil",
			args: requiredWithout{
				ThresholdUpperLimit: lo.ToPtr(7.0),
				ThresholdLowerLimit: nil,
			},
			wantErr: false,
		},
		{
			name: "both_non_nil",
			args: requiredWithout{
				ThresholdUpperLimit: lo.ToPtr(7.0),
				ThresholdLowerLimit: lo.ToPtr(6.10),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validator{}.Validate((*http.Request)(nil), tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Test_required_without() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				t.Logf("%v", err)
			}
		})
	}
}

func TestFieldError_methods(t *testing.T) {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			t.Logf("%+v", err)
		}

		t.Log("")
		for _, fe := range err.(validator.ValidationErrors) {
			t.Logf("FieldError.Tag() = %v", fe.Tag())
			t.Logf("FieldError.ActualTag() = %v", fe.ActualTag())
			t.Logf("FieldError.Namespace() = %v", fe.Namespace())
			t.Logf("FieldError.StructNamespace() = %v", fe.StructNamespace())
			t.Logf("FieldError.Field() = %v", fe.Field())
			t.Logf("FieldError.StructField() = %v", fe.StructField())
			t.Logf("FieldError.Value() = %v", fe.Value())
			t.Logf("FieldError.Param() = %v", fe.Param())
			t.Logf("FieldError.Kind() = %v", fe.Kind())
			t.Logf("FieldError.Type() = %v", fe.Type())
			t.Logf("FieldError.Error() = %v", fe.Error())
		}
	}
}
