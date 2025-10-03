package excelize_test

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExampleColumnNameToNumber() {
	fmt.Println(excelize.ColumnNameToNumber("K"))

	// Output:
	// 11 <nil>
}
