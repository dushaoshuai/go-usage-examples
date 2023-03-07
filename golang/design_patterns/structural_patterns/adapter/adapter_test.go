package adapter_test

import (
	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/structural_patterns/adapter/adapter"
)

type phone struct{}

func (p *phone) charge(adapter.V5) {}

func Example_adapter() {
	p := new(phone)
	p.charge(adapter.NewV5Adapter(adapter.NewV220()))
	// Output:
}
