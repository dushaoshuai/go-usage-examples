package Fuzz_test

import (
	"errors"
	"testing"
	"unicode/utf8"
)

func reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func TestReverse(t *testing.T) {
	for _, test := range []struct {
		in, want string
	}{
		{in: "Hello world!", want: "!dlrow olleH"},
		{in: " ", want: " "},
		{in: "!12345", want: "54321!"},
		{in: "\xec", want: "�"},
	} {
		t.Run("", func(t *testing.T) {
			rev, err := reverse(test.in)
			if err != nil {
				t.Log(err)
				t.SkipNow()
			}
			if rev != test.want {
				t.Errorf("reverse(%q) = %q, want %q", test.in, rev, test.want)
			}
		})
	}

	// $ go test -run TestReverse
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Fuzz   0.293s
}

func FuzzReverse(f *testing.F) {
	for _, s := range []string{"Hello world!", " ", "!12345"} {
		f.Add(s) // use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, origin string) {
		rev, err := reverse(origin)
		if err != nil {
			t.Log(err)
			t.SkipNow()
		}
		doubleRev, err := reverse(rev)
		if err != nil {
			t.Log(err)
			t.SkipNow()
		}
		if doubleRev != origin {
			t.Errorf("origin: %q, doubleRev: %q", origin, doubleRev)
		}
		if utf8.ValidString(origin) && !utf8.ValidString(rev) {
			t.Errorf("reverse(%q) produced invalid UTF-8 string %q", origin, rev)
		}
	})

	// fuzzing is disabled
	// $ go test -run FuzzReverse
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Fuzz   0.739s

	// fuzzing is enabled
	// $ go test -parallel 3 -fuzztime 20s -fuzz FuzzReverse
	// fuzz: elapsed: 0s, gathering baseline coverage: 0/60 completed
	// fuzz: elapsed: 0s, gathering baseline coverage: 60/60 completed, now fuzzing with 3 workers
	// fuzz: elapsed: 3s, execs: 212946 (70957/sec), new interesting: 0 (total: 60)
	// fuzz: elapsed: 6s, execs: 442558 (76537/sec), new interesting: 0 (total: 60)
	// fuzz: elapsed: 9s, execs: 671596 (76372/sec), new interesting: 0 (total: 60)
	// fuzz: elapsed: 12s, execs: 898912 (75746/sec), new interesting: 0 (total: 60)
	// fuzz: elapsed: 15s, execs: 1121910 (74356/sec), new interesting: 0 (total: 60)
	// fuzz: elapsed: 18s, execs: 1309600 (62543/sec), new interesting: 0 (total: 60)
	// fuzz: elapsed: 20s, execs: 1435549 (62180/sec), new interesting: 0 (total: 60)
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/testing/testing/Fuzz   20.524s
}
