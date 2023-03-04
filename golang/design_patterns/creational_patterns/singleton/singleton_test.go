package singleton_test

import (
	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/creational_patterns/singleton/singleton"
)

func Example_singleton() {
	singleton.GetSingleton().DoA()
	singleton.GetSingleton().DoB()
	singleton.GetSingleton().DoC()
	singleton.GetSingleton().DoD()
	singleton.GetSingleton().DoA()

	// Output:
	// Do A, my address is 0x123d380
	// Do B, my address is 0x123d380
	// Do C, my address is 0x123d380
	// Do D, my address is 0x123d380
	// Do A, my address is 0x123d380
}
