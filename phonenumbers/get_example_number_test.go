package phonenumbers

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func ExampleGetExampleNumber() {
	phoneNumber := phonenumbers.GetExampleNumber("CN")
	fmt.Println(phonenumbers.Format(phoneNumber, phonenumbers.NATIONAL))

	phoneNumber = phonenumbers.GetExampleNumberForType("CN", phonenumbers.MOBILE)
	fmt.Println(phonenumbers.Format(phoneNumber, phonenumbers.NATIONAL))

	phoneNumber = phonenumbers.GetExampleNumberForNonGeoEntity(376)
	fmt.Println(phonenumbers.Format(phoneNumber, phonenumbers.NATIONAL))

	// Output:
	// 010 1234 5678
	// 131 2345 6789
	// 0
}
