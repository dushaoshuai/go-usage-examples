package plot

import (
	_ "embed"
	"fmt"
	"image/color"
	"math/rand"
	"reflect"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
	"golang.org/x/image/font/opentype"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const (
	simsun = "simsun" // 宋体
)

var (
	//go:embed fonts/SimSun.ttf
	simsunFont []byte

	defaultLineWidth = 0.02 * vg.Inch
)

func init() {
	font.DefaultCache.Add(font.Collection{
		font.Face{
			Font: font.Font{
				Typeface: simsun,
			},
			Face: lo.Must(opentype.Parse(simsunFont)),
		},
	})
	plot.DefaultFont = font.Font{
		Typeface: simsun,
	}
}

type constraintsNumber interface {
	constraints.Integer | constraints.Float
}

func defaultTicks[T constraintsNumber](min, max T) []plot.Tick {
	defaultAxisTicks := T(7) // X/Y 轴的刻度数量
	step := (max - min) / (defaultAxisTicks - 1)

	var labelFormat string
	switch reflect.TypeOf(min).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		labelFormat = "%d"
	case reflect.Float32, reflect.Float64:
		labelFormat = "%.4f"
	}

	var ticks []plot.Tick
	for tick := min; tick <= max; tick += step {
		ticks = append(ticks, plot.Tick{
			Value: float64(tick),
			Label: fmt.Sprintf(labelFormat, tick),
		})
	}
	if T(ticks[len(ticks)-1].Value) != max {
		ticks = append(ticks, plot.Tick{
			Value: float64(max),
			Label: fmt.Sprintf(labelFormat, max),
		})
	}

	return ticks
}

// randomPoints returns some random x, y points
// with some interesting kind of trend.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func scatterExample() {
	title := "title"

	scatterData := randomPoints(50)

	scatter := lo.Must(plotter.NewScatter(scatterData))
	scatter.GlyphStyle.Color = color.RGBA{R: 27, G: 105, B: 200, A: 255} // 蓝色
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}                        // 实心圆

	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.X.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		return defaultTicks(int(min), int(max))
	})
	p.Y.Label.Text = "Y"
	p.Y.Tick.Marker = plot.TickerFunc(defaultTicks[float64])
	p.Add(plotter.NewGrid())
	p.Add(scatter)

	lo.Must0(p.Save(10*vg.Inch, 7.5*vg.Inch, "testdata/scatter.png"))
}
