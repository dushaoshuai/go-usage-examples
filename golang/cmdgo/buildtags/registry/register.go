package registry

import (
	"maps"
	"slices"
	"sync"
)

var (
	mu sync.Mutex
	m  = make(map[string]struct{})
)

func Register(pkg string) {
	mu.Lock()
	defer mu.Unlock()

	m[pkg] = struct{}{}
}

func AllPkgs() []string {
	pkgs := slices.Collect(maps.Keys(m))
	slices.Sort(pkgs)
	return pkgs
}
