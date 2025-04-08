package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

//go:generate go install
//go:generate goctl api plugin --api ./api/api.api --dir . --plugin goctl-test-plugin --style go_zero
func main() {
	p := lo.Must(plugin.NewPlugin())

	dir := lo.Must(filepath.Abs(p.Dir))
	file := filepath.Join(dir, "api.json")

	apiJSON := lo.Must(json.MarshalIndent(p, "", "  "))

	lo.Must0(os.WriteFile(file, apiJSON, 0644))
}
