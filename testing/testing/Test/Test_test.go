package Test_test

import "testing"

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("one")
	})
	t.Cleanup(func() {
		t.Log("two")
	})
	t.Cleanup(func() {
		t.Log("three")
	})

	// $ go test -test.v -test.run TestCleanup
	// === RUN   TestCleanup
	//    Test_test.go:13: three
	//    Test_test.go:10: two
	//    Test_test.go:7: one
	// --- PASS: TestCleanup (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}
