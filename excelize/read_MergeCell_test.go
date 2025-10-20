package excelize_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

// 合并区域内所有单元格的值都相同
func Test_read_MergeCell(t *testing.T) {
	f, err := excelize.OpenFile("testdata/read_MergeCell.xlsx")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	v, err := f.GetCellValue("Sheet1", "B4")
	assert.NoError(t, err)
	assert.Equal(t, v, "向前看！")

	v, err = f.GetCellValue("Sheet1", "C4")
	assert.NoError(t, err)
	assert.Equal(t, v, "向前看！")

	v, err = f.GetCellValue("Sheet1", "B5")
	assert.NoError(t, err)
	assert.Equal(t, v, "向前看！")

	v, err = f.GetCellValue("Sheet1", "C5")
	assert.NoError(t, err)
	assert.Equal(t, v, "向前看！")

	v, err = f.GetCellValue("Sheet1", "B6")
	assert.NoError(t, err)
	assert.Equal(t, v, "向前看！")

	v, err = f.GetCellValue("Sheet1", "C6")
	assert.NoError(t, err)
	assert.Equal(t, v, "向前看！")
}
