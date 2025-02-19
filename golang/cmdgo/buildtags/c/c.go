//go:build c

package c

import (
	"github.com/dushaoshuai/go-usage-examples/golang/cmdgo/buildtags/registry"
)

func init() {
	registry.Register("c")
}
