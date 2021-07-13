package viper_test

import (
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Viper uses the following precedence order. Each item takes precedence over the item below it:
// explicit call to Set
// flag
// env
// config
// key/value store
// default

func init() {
	viper.SetConfigFile("./exampleConfig.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

// Usage : go test -v -args viper_test.go --example_config none-default-config-value
//	1. don't miss "-args viper_test.go" here
//	2. package pflag implements POSIX/GNU-style flags,
//	   a single dash before an option means something different than a double dash
func ExampleBindPFlags() {
	flag.String("example_config", "default-config-value", "an example flag config")
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)
	fmt.Println(viper.GetString("example_config"))
	// Output:
	// none-default-config-value
}