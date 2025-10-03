package excelize_test

import (
	"time"

	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
)

func Example_create_document() {
	f := excelize.NewFile()
	defer f.Close()

	// create a new worksheet
	sheet2Name := "Sheet2"
	sheet2Index, err := f.NewSheet(sheet2Name)
	lo.Must0(err)

	// set value of a cell
	err = f.SetCellValue(sheet2Name, "C1", time.Now())
	if err != nil {
		panic(err)
	}
	err = f.SetCellValue(sheet2Name, "F1", 43.5654)
	if err != nil {
		panic(err)
	}

	// set the active worksheet of the workbook
	f.SetActiveSheet(sheet2Index)

	// save the spreadsheet by the given path
	lo.Must0(f.SaveAs("testdata/create_document.xlsx"))

	// Output:
}
