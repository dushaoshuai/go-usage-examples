package visitor

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/design_patterns/behavioral_patterns/visitor/visitor"
)

func Example_visitor() {
	nodes := []visitor.Node{
		&visitor.DotNode{},
		&visitor.CircleNode{},
		&visitor.RectangleNode{},
		&visitor.CompoundNode{},
	}

	visitor1 := &visitor.Visitor1{}
	for _, n := range nodes {
		visitor1.Visit(n)
	}
	fmt.Println()
	visitor2 := &visitor.Visitor2{}
	for _, n := range nodes {
		visitor2.Visit(n)
	}

	// Output:
	// visitor1 visiting a DotNode
	// visitor1 visiting a CircleNode
	// visitor1 visiting a CompoundNode
	//
	// visitor2 visiting a CircleNode
	// visitor2 visiting a RectangleNode
}
