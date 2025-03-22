package validate

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

var (
	cnmobilephonenumber = "cnmobilephonenumber"
)

func CNMobilePhoneNumber(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		parsedNumber, err := phonenumbers.Parse(field.String(), "CN")
		if err != nil {
			return false
		}
		if !phonenumbers.IsValidNumber(parsedNumber) {
			return false
		}
		return phonenumbers.GetNumberType(parsedNumber) == phonenumbers.MOBILE
	default:
		return false
	}
}
