//go:build b

package b

import (
	"github.com/dushaoshuai/go-usage-examples/golang/cmdgo/buildtags/registry"
)

func init() {
	registry.Register("b")
}
