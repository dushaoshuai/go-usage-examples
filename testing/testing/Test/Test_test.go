package Test_test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestLogFail(t *testing.T) {
	// For tests, the text will be printed only if the test fails or the -test.v flag is set.
	t.Fail()
	t.Log("one two three")

	// $ go test -test.run TestLogFail
	// --- FAIL: TestLogFail (0.00s)
	//    Test_test.go:10: one two three
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

func TestLogVFlag(t *testing.T) {
	// For tests, the text will be printed only if the test fails or the -test.v flag is set.
	t.Log("one two three")

	// $ go test -test.v -test.run TestLogVFlag
	// === RUN   TestLogVFlag
	//    Test_test.go:23: one two three
	// --- PASS: TestLogVFlag (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s

	// $ go test -test.run TestLogVFlag
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

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

func TestError(t *testing.T) {
	t.Log("one")
	// Error is equivalent to Log followed by Fail.
	t.Error("two")

	// $ go test -test.run TestError
	// --- FAIL: TestError (0.00s)
	//    Test_test.go:91: one
	//    Test_test.go:93: two
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.001s
}

func TestFailNow(t *testing.T) {
	// FailNow marks the function as having failed and stops its execution by calling runtime.Goexit.
	t.FailNow()
	t.Log("Unreachable code")

	// $ go test -test.run TestFailNow
	// --- FAIL: TestFailNow (0.00s)
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.003s
}

func TestFailed(t *testing.T) {
	if t.Failed() {
		t.Fatalf("t.Failed() = %v, want %v", t.Failed(), false)
	}
	t.Error("t.Fail()")
	if !t.Failed() {
		t.Fatalf("t.Failed() = %v, want %v", t.Failed(), true)
	}

	// $ go test -test.run TestFailed
	// --- FAIL: TestFailed (0.00s)
	//    Test_test.go:120: t.Fail()
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.001s
}

func TestFatal(t *testing.T) {
	t.Fatal("t.Fatal()")
	t.Log("Unreachable code")

	// $ go test -test.run TestFatal
	// --- FAIL: TestFatal (0.00s)
	//    Test_test.go:134: t.Fatal()
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

func TestHelper(t *testing.T) {
	t.Helper() // 暂时没看明白这个函数的作用是什么
	t.Error("t.Helper() called")

	// $ go test -test.run TestHelper
	// --- FAIL: TestHelper (0.00s)
	//    Test_test.go:147: t.Helper() called
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

func TestName(t *testing.T) {
	t.Log(t.Name())
	for i := 0; i < 2; i++ {
		t.Run("", func(t *testing.T) {
			t.Log(t.Name())
		})
	}
	for i := 0; i < 2; i++ {
		t.Run(fmt.Sprintf("subtest-%v", i), func(t *testing.T) {
			t.Log(t.Name())
		})
	}

	// $ go test -test.v -test.run TestName
	// === RUN   TestName
	//    Test_test.go:159: TestName
	// === RUN   TestName/#00
	//    Test_test.go:162: TestName/#00
	// === RUN   TestName/#01
	//    Test_test.go:162: TestName/#01
	// === RUN   TestName/subtest-0
	//    Test_test.go:167: TestName/subtest-0
	// === RUN   TestName/subtest-1
	//    Test_test.go:167: TestName/subtest-1
	// --- PASS: TestName (0.00s)
	//    --- PASS: TestName/#00 (0.00s)
	//    --- PASS: TestName/#01 (0.00s)
	//    --- PASS: TestName/subtest-0 (0.00s)
	//    --- PASS: TestName/subtest-1 (0.00s)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

func TestRun(t *testing.T) {
	for i := 0; i < 6; i++ {
		if !t.Run("", func(t *testing.T) {
			if i%2 == 0 {
				t.SkipNow() // SkipNow will not mark the test as having failed as FailNow.
			}
			t.Parallel() // Run reports whether f succeeded (or at least did not fail before calling t.Parallel).
			t.Error("failed in parallel test after calling t.Parallel")
		}) {
			t.Fatal("t.Run() = false, want true")
		}
	}

	// $ go test -test.run TestRun
	// --- FAIL: TestRun (0.00s)
	//    --- FAIL: TestRun/#01 (0.00s)
	//        Test_test.go:198: failed in parallel test after calling t.Parallel
	//    --- FAIL: TestRun/#05 (0.00s)
	//        Test_test.go:198: failed in parallel test after calling t.Parallel
	//    --- FAIL: TestRun/#03 (0.00s)
	//        Test_test.go:198: failed in parallel test after calling t.Parallel
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/testing/testing/Test   0.002s
}

func TestRun_block(t *testing.T) {
	var (
		start = time.Now()
		d     = 3 * time.Second
	)
	t.Run("block_3s", func(t *testing.T) {
		time.Sleep(d)
	})
	if time.Now().Sub(start).Round(time.Second) != d {
		t.Errorf("Current time should be 3 seconds later than the start time, but it does not.")
	}
}

func TestSetenv(t *testing.T) {
	key := "foo"
	value := "bar"
	t.Setenv(key, value)
	if got := os.Getenv(key); got != value {
		t.Errorf("os.Getenv(\"foo\") = %v, want %v", got, value)
	}

	t.Run("getenv", func(t *testing.T) {
		if got := os.Getenv(key); got != value {
			t.Errorf("os.Getenv(\"foo\") = %v, want %v", got, value)
		}
	})

	newValue := "baz"
	t.Run("setenv", func(t *testing.T) {
		// Setenv uses Cleanup to restore the environment variable to its original value after the test.
		t.Setenv(key, newValue)
		if got := os.Getenv(key); got != newValue {
			t.Errorf("os.Getenv(\"foo\") = %v, want %v", got, newValue)
		}
	})
	// After the above subtest - setenv, key's value has been restored to "bar".
	if got := os.Getenv(key); got != value {
		t.Errorf("os.Getenv(\"foo\") = %v, want %v", got, value)
	}
}

// TODO
// TempDir
// Parallel https://engineering.mercari.com/en/blog/entry/20220408-how_to_use_t_parallel/
