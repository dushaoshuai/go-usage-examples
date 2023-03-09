package cor_test

import (
	cor "github.com/dushaoshuai/go-usage-examples/golang/design_patterns/behavioral_patterns/CoR/CoR"
)

func Example_chain_of_responsibility() {
	h7 := cor.NewHandler3()

	h6 := cor.NewHandler2()
	h6.SetNext(h7)

	h5 := cor.NewHandler3()
	h5.SetNext(h6)

	h4 := cor.NewHandler1()
	h4.SetNext(h5)

	h3 := cor.NewHandler2()
	h3.SetNext(h4)

	h2 := cor.NewHandler3()
	h2.SetNext(h3)

	h1 := cor.NewHandler1()
	h1.SetNext(h2)

	h1.Do(cor.NewReq(0))

	// Output:
	// handler1, req.n = 0
	// handler3, req.n = 1
	// handler2, req.n = 2
	// handler1, req.n = 3
	// handler3, req.n = 4
	// handler2, req.n = 5
	// handler3, req.n = 6
	// handler3, req.n >= 5, stop
}
