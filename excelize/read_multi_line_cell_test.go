package excelize_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

// See testdata/read_write_multi_line_cell.png
func Test_read_write_multi_line_cell(t *testing.T) {
	file := "testdata/read_write_multi_line_cell.xlsx"
	a1Val := "hello\nworld"
	a2Val := "hello\nexcel\nhello\nword"

	f, err := excelize.OpenFile(file)
	require.Nil(t, err)
	defer func() { _ = f.Close() }()

	v, err := f.GetCellValue("Sheet1", "A1")
	require.Nil(t, err)
	require.Equal(t, a1Val, v)

	err = f.SetCellValue("Sheet1", "A2", a2Val)
	require.NoError(t, err)

	err = f.Save()
	require.NoError(t, err)

	v, err = f.GetCellValue("Sheet1", "A1")
	require.Nil(t, err)
	require.Equal(t, a1Val, v)

	v, err = f.GetCellValue("Sheet1", "A2")
	require.Nil(t, err)
	require.Equal(t, a2Val, v)
}
