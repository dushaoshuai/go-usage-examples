//go:build a

package a

import (
	"github.com/dushaoshuai/go-usage-examples/golang/cmdgo/buildtags/registry"
)

func init() {
	registry.Register("a")
}
