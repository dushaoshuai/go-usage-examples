package vibration

import (
	"math/rand"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntLine = 120
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func generateLineItemsScatter() []opts.ScatterData {
	items := make([]opts.ScatterData, 0)
	for i := 0; i < itemCntLine; i++ {
		val := rand.Intn(300)
		if i&1 == 0 {
			val = -val
		}
		items = append(items, opts.ScatterData{Value: val})
	}
	return items
}

func generateXAxis() []int {
	xAxis := make([]int, 0, itemCntLine)
	for i := 0; i <= itemCntLine; i++ {
		xAxis = append(xAxis, i)
	}
	return xAxis
}

func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

func timeDomainPlot() *charts.Scatter {
	line := charts.NewScatter()
	line.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title: "振动时域图 2024-01-06 11:16:03+08:00-2024-01-06 11:18:03+08:00",
				TitleStyle: &opts.TextStyle{
					FontWeight: "normal",
				},
				Left: "center",
				Top:  "30",
			},
		),
		charts.WithYAxisOpts(
			opts.YAxis{
				Name:         "振动",
				NameLocation: "center",
				NameGap:      35,
				Min:          -500,
				Max:          500,
			},
		),
		charts.WithXAxisOpts(
			opts.XAxis{
				SplitLine: &opts.SplitLine{
					Show: opts.Bool(true),
				},
			},
		),
		charts.WithLegendOpts(
			opts.Legend{
				Show: opts.Bool(false),
			},
		),
	)

	line.SetXAxis(generateXAxis()).
		AddSeries("振动", generateLineItemsScatter())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and label options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "markpoint options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(
				opts.MarkPointNameTypeItem{Name: "Maximum", Type: "max"},
				opts.MarkPointNameTypeItem{Name: "Average", Type: "average"},
				opts.MarkPointNameTypeItem{Name: "Minimum", Type: "min"},
			),
			charts.WithMarkPointStyleOpts(
				opts.MarkPointStyle{Label: &opts.Label{Show: opts.Bool(true)}}),
		)
	return line
}

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "splitline options",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
			},
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems(),
		charts.WithLabelOpts(
			opts.Label{Show: opts.Bool(true)},
		))
	return line
}

func lineStep() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "step style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Step: true,
			}),
		)
	return line
}

func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "smooth style",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: opts.Bool(true),
			}),
		)
	return line
}

func lineArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "area options",
		}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(
				opts.Label{
					Show: opts.Bool(true),
				}),
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.2,
				}),
		)
	return line
}

func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "smooth area"}),
	)

	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.2,
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: opts.Bool(true),
			}),
		)
	return line
}

// func lineOverlap() *charts.Line {
// 	line := charts.NewLine()
// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts"}),
// 	)
//
// 	line.SetXAxis(fruits).
// 		AddSeries("Category A", generateLineItems())
// 	line.Overlap(esEffectStyle())
// 	line.Overlap(scatterBase())
// 	return line
// }

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category  A", generateLineItems()).
		AddSeries("Category  B", generateLineItems()).
		AddSeries("Category  C", generateLineItems()).
		AddSeries("Category  D", generateLineItems())
	return line
}

func lineDemo() *charts.Line {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Search Time: Hash table vs Binary search",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Cost time(ns)",
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
			},
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Elements",
		}),
	)

	line.SetXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
		AddSeries("map", generateLineItems(),
			charts.WithLabelOpts(opts.Label{Show: opts.Bool(true), Position: "bottom"})).
		AddSeries("slice", generateLineData([]float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131}),
			charts.WithLabelOpts(opts.Label{Show: opts.Bool(true), Position: "top"})).
		SetSeriesOptions(
			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
				Name: "Average",
				Type: "average",
			}),
			charts.WithLineChartOpts(opts.LineChart{
				Smooth: opts.Bool(true),
			}),
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show:      opts.Bool(true),
					Formatter: "{a}: {b}",
				},
			}),
		)

	return line
}

func lineSymbols() *charts.Line {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "symbol options",
			Subtitle: "tooltip with 'axis' trigger",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}),
	)

	// Put data into instance
	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}
