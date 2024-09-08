package simple_but_lower_level_read_test

import (
	"fmt"
	"log/slog"

	parquetfile "github.com/apache/arrow/go/v17/parquet/file"
)

func Example_simple_read() {
	rdr, err := parquetfile.OpenParquetFile("../../../../../../../example.parquet", true)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer rdr.Close()

	for i := range rdr.NumRowGroups() {
		rgr := rdr.RowGroup(i)

		const colwidth = 17

		scanners := make([]*Dumper, rgr.NumColumns())
		for c := range rgr.NumColumns() {
			col, err := rgr.Column(c)
			if err != nil {
				slog.Error("unable to fetch column",
					slog.Int("column", c),
					slog.Any("err", err),
				)
			}
			scanners[c] = createDumper(col)
			fmt.Printf(fmt.Sprintf("%%-%ds|", colwidth), col.Descriptor().Name())
		}
		fmt.Println()

		var line string
		for {
			data := false
			for _, s := range scanners {
				if val, ok := s.Next(); ok {
					if !data {
						fmt.Print(line)
					}
					fmt.Print(s.FormatValue(val, colwidth), "|")
					data = true
				} else {
					if data {
						fmt.Printf(fmt.Sprintf("%%-%ds|", colwidth), "")
					} else {
						line += fmt.Sprintf(fmt.Sprintf("%%-%ds|", colwidth), "")
					}
				}
			}
			if !data {
				break
			}
			fmt.Println()
			line = ""
		}
		fmt.Println()
	}

	// Output:
	// carat            |cut              |color            |clarity          |depth            |table            |price            |x                |y                |z                |__index_level_0__|
	// 0.230000         |Ideal            |E                |SI2              |61.500000        |55.000000        |326              |3.950000         |3.980000         |2.430000         |0                |
	// 0.210000         |Premium          |E                |SI1              |59.800000        |61.000000        |326              |3.890000         |3.840000         |2.310000         |1                |
	// 0.230000         |Good             |E                |VS1              |56.900000        |65.000000        |327              |4.050000         |4.070000         |2.310000         |2                |
	// 0.290000         |Premium          |I                |VS2              |62.400000        |58.000000        |334              |4.200000         |4.230000         |2.630000         |3                |
	// 0.310000         |Good             |J                |SI2              |63.300000        |58.000000        |335              |4.340000         |4.350000         |2.750000         |4                |
	// 0.240000         |Very Good        |J                |VVS2             |62.800000        |57.000000        |336              |3.940000         |3.960000         |2.480000         |5                |
	// 0.240000         |Very Good        |I                |VVS1             |62.300000        |57.000000        |336              |3.950000         |3.980000         |2.470000         |6                |
	// 0.260000         |Very Good        |H                |SI1              |61.900000        |55.000000        |337              |4.070000         |4.110000         |2.530000         |7                |
	// 0.220000         |Fair             |E                |VS2              |65.100000        |61.000000        |337              |3.870000         |3.780000         |2.490000         |8                |
	// 0.230000         |Very Good        |H                |VS1              |59.400000        |61.000000        |338              |4.000000         |4.050000         |2.390000         |9                |
}
