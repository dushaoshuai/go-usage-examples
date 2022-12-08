package phonenumbers_test

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func ExampleFormat() {
	phoneNumber, err := phonenumbers.Parse("17629343048", "CN")
	if err != nil {
		panic(err)
	}
	if !phonenumbers.IsValidNumber(phoneNumber) {
		panic("invalid number")
	}

	formats := []phonenumbers.PhoneNumberFormat{
		phonenumbers.E164,
		phonenumbers.INTERNATIONAL,
		phonenumbers.NATIONAL,
		phonenumbers.RFC3966,
	}
	for _, f := range formats {
		fmt.Println(phonenumbers.Format(phoneNumber, f))
	}

	// Output:
	// +8617629343048
	// +86 176 2934 3048
	// 176 2934 3048
	// tel:+86-176-2934-3048
}
