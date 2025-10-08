package excelize_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func TestSetCellValue_GetCellType(t *testing.T) {
	tfpath := filepath.Join("testdata", "SetCellValue_GetCellType.xlsx")
	require.NoError(t, os.RemoveAll(tfpath))

	file := excelize.NewFile()
	defer file.Close()
	defer file.SaveAs(tfpath)

	require.NoError(t, file.SetCellValue("Sheet1", "A1", 1))
	require.NoError(t, file.SetCellValue("Sheet1", "A2", int8(2)))
	require.NoError(t, file.SetCellValue("Sheet1", "A3", int16(3)))
	require.NoError(t, file.SetCellValue("Sheet1", "A4", int32(4)))
	require.NoError(t, file.SetCellValue("Sheet1", "A5", int64(5)))
	require.NoError(t, file.SetCellValue("Sheet1", "A6", uint(6)))
	require.NoError(t, file.SetCellValue("Sheet1", "A7", uint8(7)))
	require.NoError(t, file.SetCellValue("Sheet1", "A8", uint16(8)))
	require.NoError(t, file.SetCellValue("Sheet1", "A9", uint32(9)))
	require.NoError(t, file.SetCellValue("Sheet1", "A10", uint64(10)))
	require.NoError(t, file.SetCellValue("Sheet1", "A11", float32(11.17323)))
	require.NoError(t, file.SetCellValue("Sheet1", "A12", float64(12.17323)))
	require.NoError(t, file.SetCellValue("Sheet1", "A13", "13 in A13"))         // SharedString
	require.NoError(t, file.SetCellValue("Sheet1", "A14", []byte("14 in A14"))) // SharedString
	require.NoError(t, file.SetCellValue("Sheet1", "A15", 15*time.Second))
	require.NoError(t, file.SetCellValue("Sheet1", "A16", time.Now()))
	require.NoError(t, file.SetCellValue("Sheet1", "A17", true))          // Bool
	require.NoError(t, file.SetCellValue("Sheet1", "A18", false))         // Bool
	require.NoError(t, file.SetCellBool("Sheet1", "A19", true))           // Bool
	require.NoError(t, file.SetCellBool("Sheet1", "A20", false))          // Bool
	require.NoError(t, file.SetCellDefault("Sheet1", "A21", "21 in A21")) // InlineString
	require.NoError(t, file.SetCellInt("Sheet1", "A22", 22))
	require.NoError(t, file.SetCellUint("Sheet1", "A23", 23))
	require.NoError(t, file.SetCellFloat("Sheet1", "A24", 24.17323, 6, 64))
	require.NoError(t, file.SetCellStr("Sheet1", "A25", "25 in A25"))          // SharedString
	require.NoError(t, file.SetCellFormula("Sheet1", "A26", "=SUM(A1,A2,A3)")) // Formula (normal formula)

	require.NoError(t, file.SetSheetDimension("Sheet1", "A1:A26"))

	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 4, 0, '\t', 0)
	defer tabwriter.Flush()

	dimension, err := file.GetSheetDimension("Sheet1")
	require.NoError(t, err)

	// A1:A27 -> A27 -> 27
	_, maxRows, err := excelize.SplitCellName(strings.Split(dimension, ":")[1])
	require.NoError(t, err)

	for r := range maxRows {
		cellName, err := excelize.JoinCellName("A", r+1)
		require.NoError(t, err)

		cellType, err := file.GetCellType("Sheet1", cellName)
		require.NoError(t, err)

		cellValue, err := file.GetCellValue("Sheet1", cellName)
		require.NoError(t, err)

		rawCellValue, err := file.GetCellValue("Sheet1", cellName, excelize.Options{RawCellValue: true})
		require.NoError(t, err)

		if cellType == excelize.CellTypeFormula {
			cellValue, err = file.CalcCellValue("Sheet1", cellName)
		}

		fmt.Fprintf(tabwriter, "%s\t%v\t%s\t%s\t\n", cellName, CellType(cellType).String(), cellValue, rawCellValue)
	}

	// Output:
	// A1	Unset		1				1
	// A2	Unset		2				2
	// A3	Unset		3				3
	// A4	Unset		4				4
	// A5	Unset		5				5
	// A6	Unset		6				6
	// A7	Unset		7				7
	// A8	Unset		8				8
	// A9	Unset		9				9
	// A10	Unset		10				10
	// A11	Unset		11.17323		11.17323
	// A12	Unset		12.17323		12.17323
	// A13	SharedString13 in A13		13 in A13
	// A14	SharedString14 in A14		14 in A14
	// A15	Unset		00:00:15		0.00017361112
	// A16	Unset		10/8/25 17:40	45938.7364057145
	// A17	Bool		TRUE			1
	// A18	Bool		FALSE			0
	// A19	Bool		TRUE			1
	// A20	Bool		FALSE			0
	// A21	InlineString21 in A21		21 in A21
	// A22	Unset		22				22
	// A23	Unset		23				23
	// A24	Unset		24.17323		24.17323
	// A25	SharedString25 in A25		25 in A25
	// A26	Formula		6
}
