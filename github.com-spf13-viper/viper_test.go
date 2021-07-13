package viper_test

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Viper uses the following precedence order. Each item takes precedence over the item below it:
// explicit call to Set
// flag
// env
// config file
// key/value store
// default

// Usage : go test -v -args viper_test.go --example_config none-default-config-value
//	1. don't miss "-args viper_test.go" here
//	2. package pflag implements POSIX/GNU-style flags,
//	   a single dash before an option means something different than a double dash
func ExampleBindPFlags() {
	pflag.String("example_config", "default-config-value", "an example flag config")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Println(viper.GetString("example_config"))
	// Output:
	// none-default-config-value
}
