package vibration

import (
	"io"
	"os"
	"testing"

	"github.com/go-echarts/go-echarts/v2/components"
)

func Test_line(t *testing.T) {
	page := components.NewPage()
	page.AddCharts(
		timeDomainPlot(),
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
