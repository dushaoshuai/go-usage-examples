package phonenumbers_test

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func ExampleParseWithDefaultRegion() {
	numbers := []string{
		"+8617823456789",
		"17823456789",
		"+862067799192",
		"+12067799192",
		"2067799192",
	}

	for _, number := range numbers {
		// Parse 的第二个参数 defaultRegion 只有在第一个参数 numberToParse 不包含 country calling code 时才起作用，
		// 在这种情况下，numberToParse 会被解析为 defaultRegion 内的电话号码。
		parsedNumber, err := phonenumbers.Parse(number, "CN")
		if err != nil {
			fmt.Println(err, parsedNumber)
			continue
		}
		fmt.Println(
			number,
			parsedNumber.GetCountryCode(),
			phonenumbers.GetRegionCodeForNumber(parsedNumber),
			phonenumbers.IsValidNumber(parsedNumber),
		)
	}

	// Output:
	// +8617823456789 86 CN true
	// 17823456789 86 CN true
	// +862067799192 86 CN true
	// +12067799192 1 US true
	// 2067799192 86 CN true
}

func ExampleParseWithoutDefaultRegion() {
	numbers := []string{
		"17823456789",
		"2067799192",
	}

	for _, number := range numbers {
		_, err := phonenumbers.Parse(number, "")
		fmt.Println(number, err)
	}

	// Output:
	// 17823456789 invalid country code
	// 2067799192 invalid country code
}
