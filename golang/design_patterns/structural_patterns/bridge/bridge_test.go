package bridge_test

import (
	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/structural_patterns/bridge/bridge"
)

func Example_bridge() {
	linux := bridge.NewLinux()
	linux.SetDisplay(bridge.NewAoc())
	linux.Show()
	linux.SetDisplay(bridge.NewBenq())
	linux.Show()

	mac := bridge.NewMac()
	mac.SetDisplay(bridge.NewBenq())
	mac.Show()
	mac.SetDisplay(bridge.NewAoc())
	mac.Show()

	// Output:
}
