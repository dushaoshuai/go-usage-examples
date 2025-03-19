package ristretto

import (
	"strconv"
	"testing"
)

func Test_newCache(t *testing.T) {
	cache, err := newCache()
	if err != nil {
		t.Fatal(err)
	}

	for i := range maxCost {
		if !cache.Set(strconv.FormatInt(i, 10), i, 1) {
			t.Fatalf("set key failed: %d", i)
		}
		cache.Wait()
	}

	for i := range maxCost {
		_, ok := cache.Get(strconv.FormatInt(i, 10))
		if !ok {
			t.Errorf("key not found in cache: %d", i)
		}
	}

	cache.Set("one_more", 77, 1)
	cache.Wait()

	for i := range maxCost {
		_, ok := cache.Get(strconv.FormatInt(i, 10))
		if !ok {
			t.Logf("key not found in cache: %d", i)
		}
	}
	_, ok := cache.Get("one_more")
	if !ok {
		t.Logf("key not found in cache: %s", "one_more")
	}

	// Output:
	// key not found in cache: 2522
}
