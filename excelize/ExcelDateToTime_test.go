package excelize_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func Test_ExcelDateToTime(t *testing.T) {
	file := "testdata/ExcelDateToTime_test.xlsx"

	f, err := excelize.OpenFile(file)
	require.Nil(t, err)
	defer func() { _ = f.Close() }()

	// 1/2/2022
	v, err := f.GetCellValue("Sheet1", "A1")
	require.Nil(t, err)

	float, err := strconv.ParseFloat(v, 64)
	require.NoError(t, err)

	// Windows 版 Excel（绝大多数情况）	1900
	// macOS 版 Excel（旧版或特定模板）	1904
	// LibreOffice / WPS	1900
	// Google Sheets	1900（与 Excel Windows 兼容）
	tim, err := excelize.ExcelDateToTime(float, false) // 默认时区是 UTC
	require.NoError(t, err)

	expected := time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC)
	require.True(t, tim.Equal(expected), "expected %v, got %v", expected, tim)
}
