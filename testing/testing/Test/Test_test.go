package Test_test

import (
	"testing"
	"time"
)

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

	// Log doc: For tests, the text will be printed only if the test fails or the -test.v flag is set.
	// $ go test -test.v -test.run TestCleanup
	// === RUN   TestCleanup
	//    Test_test.go:13: three
	//    Test_test.go:10: two
	//    Test_test.go:7: one
	// --- PASS: TestCleanup (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

func TestDeadline(t *testing.T) {
	t.Log(time.Now())
	t.Log(t.Deadline())

	// The default timeout is 10 minutes.
	// $ go test -test.v -test.run TestDeadline
	// === RUN   TestDeadline
	//    Test_test.go:31: 2023-04-11 09:04:59.378734802 +0800 CST m=+0.000155140
	//    Test_test.go:32: 2023-04-11 09:14:59.37866619 +0800 CST m=+600.000086528 true
	// --- PASS: TestDeadline (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.001s

	// $ go test -test.v -test.timeout 1m -test.run TestDeadline
	// === RUN   TestDeadline
	//    Test_test.go:31: 2023-04-11 09:03:18.886912617 +0800 CST m=+0.000173185
	//    Test_test.go:32: 2023-04-11 09:04:18.886832252 +0800 CST m=+60.000092820 true
	// --- PASS: TestDeadline (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.001s

	// "-timeout 0" disable the timeout.
	// $ go test -test.v -test.timeout 0 -test.run TestDeadline
	// === RUN   TestDeadline
	//    Test_test.go:31: 2023-04-11 09:05:45.317425146 +0800 CST m=+0.000283747
	//    Test_test.go:32: 0001-01-01 00:00:00 +0000 UTC false
	// --- PASS: TestDeadline (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}
