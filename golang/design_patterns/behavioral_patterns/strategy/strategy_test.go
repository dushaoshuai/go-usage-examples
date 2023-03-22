package strategy_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/behavioral_patterns/strategy/strategy"
)

func ExampleStrategy() {
	c := strategy.NewContext(&strategy.Add{})
	fmt.Println(c.ExecuteStrategy(3, 4))

	action := "subtraction" // user's desired action
	switch action {
	case "addition":
		c.SetStrategy(&strategy.Add{})
	case "subtraction":
		c.SetStrategy(&strategy.Subtract{})
	case "multiplication":
		c.SetStrategy(&strategy.Multiply{})
	}
	fmt.Println(c.ExecuteStrategy(4, 5))

	// Output:
	// 7
	// -1
}
