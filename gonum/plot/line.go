package plot

import (
	"image/color"
	"slices"

	"github.com/samber/lo"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func lineExample() {
	title := "title"

	lineData := randomPoints(50)

	line := lo.Must(plotter.NewLine(lineData))
	line.LineStyle.Color = color.RGBA{R: 27, G: 105, B: 200, A: 255} // 蓝色
	line.LineStyle.Width = defaultLineWidth

	p := plot.New()
	p.Title.Text = title
	p.Add(plotter.NewGrid())

	p.Legend.Padding = 0.15 * vg.Inch
	p.Legend.YOffs = 0.3 * vg.Inch
	p.Legend.XOffs = -0.3 * vg.Inch
	p.Legend.ThumbnailWidth = 0.63 * vg.Inch

	p.X.Label.Text = "X"
	p.X.Label.Padding = 0.1 * vg.Inch
	p.X.Tick.Marker = plot.TickerFunc(defaultTicks[float64])
	p.Y.Label.Text = "Y"
	p.Y.Tick.Marker = plot.TickerFunc(defaultTicks[float64])
	p.Add(line)

	// markline
	for markLine := range slices.Values([]struct {
		name  string
		xVal  float64
		style draw.LineStyle
	}{
		{
			name: "xMax/4",
			xVal: p.X.Max / 4,
			style: draw.LineStyle{
				Color: color.RGBA{R: 255, G: 0, B: 0, A: 255},
				Width: defaultLineWidth,
			},
		},
		{
			name: "xMax/2.5",
			xVal: p.X.Max / 2.5,
			style: draw.LineStyle{
				Color: color.RGBA{R: 202, G: 202, B: 44, A: 255},
				Width: defaultLineWidth,
			},
		},
		{
			name: "xMax/2",
			xVal: p.X.Max / 2,
			style: draw.LineStyle{
				Color:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
				Width:  defaultLineWidth,
				Dashes: []vg.Length{0.12 * vg.Inch, 0.04 * vg.Inch},
			},
		},
		{
			name: "xMax/1.2",
			xVal: p.X.Max / 1.2,
			style: draw.LineStyle{
				Color:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
				Width:  defaultLineWidth,
				Dashes: []vg.Length{0.3 * vg.Inch, 0.04 * vg.Inch},
			},
		},
	}) {
		mLine := lo.Must(plotter.NewLine(plotter.XYs{
			{markLine.xVal, 0},
			{markLine.xVal, p.Y.Max},
		}))
		mLine.LineStyle = markLine.style
		p.Add(mLine)
		p.Legend.Add(markLine.name, mLine)
	}

	lo.Must0(p.Save(10*vg.Inch, 7.5*vg.Inch, "testdata/line.png"))
}
