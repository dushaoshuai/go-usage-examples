package phonenumbers_test

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func ExamplePhoneNumberType() {
	// 填移动号码
	fmt.Println(getNumberType("xxxxxxxx"))
	// 填座机号码
	fmt.Println(getNumberType("02081167888"))

	// Output:
	// 1
	// 0
}

func getNumberType(numberToParse string) phonenumbers.PhoneNumberType {
	number, err := phonenumbers.Parse(numberToParse, "CN")
	if err != nil {
		panic(err)
	}
	return phonenumbers.GetNumberType(number)
}
