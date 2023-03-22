package visitor

import (
	"fmt"
)

type Node interface {
	IsNode()
}

type DotNode struct{}

func (d *DotNode) IsNode() {}

type CircleNode struct{}

func (c *CircleNode) IsNode() {}

type RectangleNode struct{}

func (r *RectangleNode) IsNode() {}

type CompoundNode struct{}

func (c *CompoundNode) IsNode() {}

type Visitor interface {
	Visit(Node)
}

type Visitor1 struct{}

func (v *Visitor1) Visit(n Node) {
	switch n.(type) { // a Visitor doesn't have to visit all Node type
	case *DotNode:
		fmt.Println("visitor1 visiting a DotNode")
	case *CircleNode:
		fmt.Println("visitor1 visiting a CircleNode")
	case *CompoundNode:
		fmt.Println("visitor1 visiting a CompoundNode")
	}
}

type Visitor2 struct{}

func (v *Visitor2) Visit(n Node) {
	switch n.(type) { // a Visitor doesn't have to visit all Node type
	case *CircleNode:
		fmt.Println("visitor2 visiting a CircleNode")
	case *RectangleNode:
		fmt.Println("visitor2 visiting a RectangleNode")
	}
}
