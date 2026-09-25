package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func Test_AddGet(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cache.Add(c.key, c.val)
			actual, ok := cache.Get(c.key)

			if !ok {
				t.Fatalf("expected to find key %q", c.key)
			}

			if !bytes.Equal(actual, c.val) {
				t.Errorf("want %q as value, got %q instead", c.val, actual)
			}
		})
	}
}

func Test_reapLoop(t *testing.T) {
	interval := 5 * time.Second
	waitTime := interval * 2
	cache := NewCache(interval)
	testCase := struct {
		key string
		val []byte
	}{
		key: "https://example.com",
		val: []byte("testdata"),
	}

	cache.Add(testCase.key, testCase.val)

	if _, ok := cache.Get(testCase.key); !ok {
		t.Fatalf("expected to find key %q", testCase.key)
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get(testCase.key); ok {
		t.Errorf("expected to not find key %q", testCase.key)
	}
}
