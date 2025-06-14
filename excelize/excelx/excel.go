package excelx

import (
	"fmt"
	"iter"
	"slices"

	"github.com/xuri/excelize/v2"
)

type SheetSelector func(f *excelize.File) []string

func NewAllSheetsSelector() func(f *excelize.File) []string {
	return func(f *excelize.File) []string {
		return f.GetSheetList()
	}
}

func NewActiveSheetSelector() func(f *excelize.File) []string {
	return func(f *excelize.File) []string {
		return []string{
			f.GetSheetName(f.GetActiveSheetIndex()),
		}
	}
}

func NewNamedSheetSelector(sheets ...string) func(f *excelize.File) []string {
	return func(_ *excelize.File) []string {
		return sheets
	}
}

func ReadFile(filePath string, sheetSelector SheetSelector) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Read(f, sheetSelector)
}

func Read(f *excelize.File, sheetSelector SheetSelector) ([][]string, error) {
	var result [][]string

	for sheet := range slices.Values(sheetSelector(f)) {
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, err
		}
		result = append(result, rows...)
	}

	return result, nil
}

func WriteSheet(f *excelize.File, sheetName string, header []string, rows iter.Seq[[]string]) error {
	var rowNum int
	getStartingCell := func() string {
		rowNum++
		return fmt.Sprintf("A%d", rowNum)
	}

	// set header
	err := f.SetSheetRow(sheetName, getStartingCell(), &header)
	if err != nil {
		return err
	}

	// set rows
	for row := range rows {
		err = f.SetSheetRow(sheetName, getStartingCell(), &row)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteActiveSheet(f *excelize.File, header []string, rows iter.Seq[[]string]) error {
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
	return WriteSheet(f, sheetName, header, rows)
}
