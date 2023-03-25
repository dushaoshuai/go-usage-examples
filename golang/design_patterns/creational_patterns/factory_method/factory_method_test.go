package factory_method_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/creational_patterns/factory_method/factory_method"
)

func Example_factory_method() {
	echo := factory_method.NewEcho()
	fmt.Println(echo.Name(), echo.Execute())

	cd := factory_method.NewCd()
	fmt.Println(cd.Name(), cd.Execute())

	// Output:
	// echo <nil>
	// cd <nil>
}
