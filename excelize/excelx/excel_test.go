package excelx

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

var (
	testingSheet1Name = "Sheet1"
	testingSheet1Data = [][]string{
		{"1A", "1B", "1C", "1D", "", "1F", "", "", "1I", "1J"},
		{"2A", "", "", "", "", "", "", "", "2I"},
		nil,
		{"", "4B"},
		nil,
		{"", "", "", "6D", "6E", "6F", "6G", "6H", "6I"},
		nil,
		{"8A"},
	}

	testingSheet2Name = "Sheet2"
	testingSheet2Data = [][]string{
		nil,
		nil,
		nil,
		nil,
		nil,
		{"", "", "6C"},
	}

	testingSheet3Name = "测试Sheet3"
	testingSheet3Data = [][]string{
		{"1A"},
	}
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name          string
		sheetSelector SheetSelector
		want          [][]string
		wantErr       bool
	}{
		{
			name:          "all sheets",
			sheetSelector: NewAllSheetsSelector(),
			want:          slices.Concat(testingSheet1Data, testingSheet2Data, testingSheet3Data),
			wantErr:       false,
		},
		{
			name:          "active sheet",
			sheetSelector: NewActiveSheetSelector(),
			want:          testingSheet1Data,
			wantErr:       false,
		},
		{
			name:          testingSheet1Name,
			sheetSelector: NewNamedSheetSelector(testingSheet1Name),
			want:          testingSheet1Data,
			wantErr:       false,
		},
		{
			name:          testingSheet2Name,
			sheetSelector: NewNamedSheetSelector(testingSheet2Name),
			want:          testingSheet2Data,
			wantErr:       false,
		},
		{
			name:          testingSheet3Name,
			sheetSelector: NewNamedSheetSelector(testingSheet3Name),
			want:          testingSheet3Data,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile("./testdata/testing.xlsx", tt.sheetSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		name          string
		sheetSelector SheetSelector
		want          [][]string
		wantErr       bool
	}{
		{
			name:          "all sheets",
			sheetSelector: NewAllSheetsSelector(),
			want:          slices.Concat(testingSheet1Data, testingSheet2Data, testingSheet3Data),
			wantErr:       false,
		},
		{
			name:          "active sheet",
			sheetSelector: NewActiveSheetSelector(),
			want:          testingSheet1Data,
			wantErr:       false,
		},
		{
			name:          testingSheet1Name,
			sheetSelector: NewNamedSheetSelector(testingSheet1Name),
			want:          testingSheet1Data,
			wantErr:       false,
		},
		{
			name:          testingSheet2Name,
			sheetSelector: NewNamedSheetSelector(testingSheet2Name),
			want:          testingSheet2Data,
			wantErr:       false,
		},
		{
			name:          testingSheet3Name,
			sheetSelector: NewNamedSheetSelector(testingSheet3Name),
			want:          testingSheet3Data,
			wantErr:       false,
		},
	}

	f, err := excelize.OpenFile("./testdata/testing.xlsx")
	require.NoError(t, err)
	defer f.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(f, tt.sheetSelector)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteSheet(t *testing.T) {
	f := excelize.NewFile()

	header := []string{
		"header 1A", "header 1B", "", "", "header 1E",
	}

	err := WriteSheet(f, testingSheet1Name, header, slices.Values(testingSheet1Data))
	require.NoError(t, err)

	got, err := Read(f, NewNamedSheetSelector(testingSheet1Name))
	require.NoError(t, err)

	require.Equal(t, slices.Concat([][]string{header}, testingSheet1Data), got)
}

func TestWriteActiveSheet(t *testing.T) {
	f := excelize.NewFile()

	header := []string{
		"header 1A", "header 1B", "", "", "header 1E",
	}

	err := WriteActiveSheet(f, header, slices.Values(testingSheet1Data))
	require.NoError(t, err)

	got, err := Read(f, NewActiveSheetSelector())
	require.NoError(t, err)

	require.Equal(t, slices.Concat([][]string{header}, testingSheet1Data), got)
}
