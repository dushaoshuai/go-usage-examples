package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

const (
	sheet1 = "Sheet1"
	sheet2 = "Sheet2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	f.NewSheet(sheet2)
	// Set value of a cell.
	err := f.SetCellValue(sheet2, "A2", "Hello world")
	if err != nil {
		logrus.Fatal(err)
	}
	err = f.SetCellValue(sheet1, "B2", 100)
	if err != nil {
		logrus.Fatal(err)
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(f.GetSheetIndex(sheet1))
	// Add comment.
	err = f.AddComment(sheet1, "B2", `{"author":"杜少帅: ","text":"This is a comment."}`)
	if err != nil {
		logrus.Fatal(err)
	}
	// Save spreadsheet by the given path.
	err = f.SaveAs("Book1.xlsx")
	if err != nil {
		logrus.Fatal(err)
	}
}
