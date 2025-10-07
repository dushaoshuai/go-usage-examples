package excelize_test

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

//go:generate go run golang.org/x/tools/cmd/stringer -output consts_string_test.go -type CellType -trimprefix CellType
type CellType byte

// Cell value types enumeration.
const (
	CellTypeUnset CellType = iota
	CellTypeBool
	CellTypeDate
	CellTypeError
	CellTypeFormula
	CellTypeInlineString
	CellTypeNumber
	CellTypeSharedString
)

func TestGetCellType(t *testing.T) {
	file, err := excelize.OpenFile("testdata/GetCellType.xlsx")
	require.NoError(t, err)
	defer file.Close()

	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 0, '\t', 0)
	defer writer.Flush()

	for r := range 24 {
		cellName, err := excelize.JoinCellName("A", r+1)
		require.NoError(t, err)

		cellType, err := file.GetCellType("Sheet1", cellName)
		require.NoError(t, err)

		cellValue, err := file.GetCellValue("Sheet1", cellName)
		require.NoError(t, err)

		rawCellValue, err := file.GetCellValue("Sheet1", cellName, excelize.Options{RawCellValue: true})
		require.NoError(t, err)

		fmt.Fprintf(writer, "%s\t%v\t%s\t%s\t\n", cellName, CellType(cellType).String(), cellValue, rawCellValue)
	}

	// Output:
	// A1	Unset		1						1
	// A2	Unset		2.00					2
	// A3	Unset		$3.00					3
	// A4	Unset		 $4.00 					4
	// A5	Unset		Thursday,January 4,1900	5
	// A6	Unset		0:00:00					6
	// A7	Unset		700.00%					7
	// A8	Unset		8    					8
	// A9	Unset		9.00E+00				9
	// A10	Unset		10						10
	// A11	Unset		00011					11
	// A12	Unset		00000-0012				12
	// A13	Unset		000-00-0013				13
	// A14	SharedString我						我
	// A15	SharedString你						你
	// A16	SharedString他						他
	// A17	Bool		TRUE					1
	// A18	Bool		FALSE					0
	// A19	Unset		10/6/25 10:12 AM		45936.4254050926
	// A20	Unset		45937					45937
	// A21	Formula		9/27					9/27
	// A22	Unset		1						1
	// A23	Unset		3						3
	// A24	Unset		10/7/2025				45937
}
