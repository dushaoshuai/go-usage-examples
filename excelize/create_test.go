package excelize_test

import "github.com/xuri/excelize/v2"

func Example_create_document() {

	file := excelize.NewFile()

	sheetName := "xiaoying"
	sheet2 := file.NewSheet(sheetName)

	_ = file.SetCellValue(sheetName, "E11", "xiaoying")
	_ = file.SetCellValue(sheetName, "F11", "ying")

	_ = file.SetCellValue(sheetName, "E12", "shao")
	_ = file.SetCellValue(sheetName, "F12", "shuai")

	file.SetActiveSheet(sheet2)

	if err := file.SaveAs("xiaoyingHeShaoshuai.xlsx"); err != nil {
		panic(err)
	}

	// Output:
}
