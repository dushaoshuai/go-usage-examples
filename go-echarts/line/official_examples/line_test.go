package official_examples

import (
	"io"
	"os"
	"testing"

	"github.com/go-echarts/go-echarts/v2/components"
)

func Test_line(t *testing.T) {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		lineShowLabel(),
		lineSymbols(),
		lineMarkPoint(),
		lineSplitLine(),
		lineStep(),
		lineSmooth(),
		lineArea(),
		lineSmoothArea(),
		// lineOverlap(),
		lineMulti(),
		lineDemo(),
	)
	f, err := os.Create("line.html")
	if err != nil {
		t.Fatal(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		t.Fatal(err)
	}
}
