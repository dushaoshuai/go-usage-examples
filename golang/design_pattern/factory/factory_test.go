package factory

import (
	"fmt"

	"api-examples/golang/design_pattern/factory/interface_factory"
)

func Example_interface_factory() {
	echo := interface_factory.NewEcho()
	fmt.Println(echo.Name(), echo.Execute())

	cd := interface_factory.NewCd()
	fmt.Println(cd.Name(), cd.Execute())

	// Output:
	// echo <nil>
	// cd <nil>
}
