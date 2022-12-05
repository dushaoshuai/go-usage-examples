package phonenumbers

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func ExampleFormat() {
	phoneNumber, err := phonenumbers.Parse("18829352048", "CN")
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
	// +8618829352048
	// +86 188 2935 2048
	// 188 2935 2048
	// tel:+86-188-2935-2048
}
