package parquet_go_test

import (
	"fmt"
	"log/slog"

	"github.com/parquet-go/parquet-go"
)

func Example_simple_read() {
	type Row struct {
		Carat           float64 `parquet:"carat"`
		Cut             string  `parquet:"cut"`
		Color           string  `parquet:"color"`
		Clarity         string  `parquet:"clarity"`
		Depth           float64 `parquet:"depth"`
		Table           float64 `parquet:"table"`
		Price           int     `parquet:"price"`
		X               float64 `parquet:"x"`
		Y               float64 `parquet:"y"`
		Z               float64 `parquet:"z"`
		Index_level_0__ int     `parquet:"__index_level_0__"`
	}

	rows, err := parquet.ReadFile[Row]("../../../example.parquet")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for _, row := range rows {
		fmt.Printf("%+v\n", row)
	}

	// Output:
	// {Carat:0.23 Cut:Ideal Color:E Clarity:SI2 Depth:61.5 Table:55 Price:326 X:3.95 Y:3.98 Z:2.43 Index_level_0__:0}
	// {Carat:0.21 Cut:Premium Color:E Clarity:SI1 Depth:59.8 Table:61 Price:326 X:3.89 Y:3.84 Z:2.31 Index_level_0__:1}
	// {Carat:0.23 Cut:Good Color:E Clarity:VS1 Depth:56.9 Table:65 Price:327 X:4.05 Y:4.07 Z:2.31 Index_level_0__:2}
	// {Carat:0.29 Cut:Premium Color:I Clarity:VS2 Depth:62.4 Table:58 Price:334 X:4.2 Y:4.23 Z:2.63 Index_level_0__:3}
	// {Carat:0.31 Cut:Good Color:J Clarity:SI2 Depth:63.3 Table:58 Price:335 X:4.34 Y:4.35 Z:2.75 Index_level_0__:4}
	// {Carat:0.24 Cut:Very Good Color:J Clarity:VVS2 Depth:62.8 Table:57 Price:336 X:3.94 Y:3.96 Z:2.48 Index_level_0__:5}
	// {Carat:0.24 Cut:Very Good Color:I Clarity:VVS1 Depth:62.3 Table:57 Price:336 X:3.95 Y:3.98 Z:2.47 Index_level_0__:6}
	// {Carat:0.26 Cut:Very Good Color:H Clarity:SI1 Depth:61.9 Table:55 Price:337 X:4.07 Y:4.11 Z:2.53 Index_level_0__:7}
	// {Carat:0.22 Cut:Fair Color:E Clarity:VS2 Depth:65.1 Table:61 Price:337 X:3.87 Y:3.78 Z:2.49 Index_level_0__:8}
	// {Carat:0.23 Cut:Very Good Color:H Clarity:VS1 Depth:59.4 Table:61 Price:338 X:4 Y:4.05 Z:2.39 Index_level_0__:9}
}
