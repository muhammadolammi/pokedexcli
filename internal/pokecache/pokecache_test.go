package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestPokeCache(t *testing.T) {
	const interval = 5 * time.Second

	cache := NewCache(interval)
	// test if map is not empty
	if cache.cacheMap == nil {
		t.Error("the cache map is nil")
	}

}

func TestAddGetPokeCache(t *testing.T) {
	const interval = 10 * time.Second

	// test if map is not empty
	cases := []struct {
		key string
		val []byte
	}{
		{
			"key1",
			[]byte("val1"),
		},
		{
			"key2",
			[]byte("val2"),
		},
		{
			"key3",
			[]byte("val3"),
		},
		{
			"key4",
			[]byte("val4"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}

}

func TestPokeCacheReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("faith", []byte("is an amazing human"))
	v, ok := cache.Get("faith")
	if !ok {
		t.Errorf("%s shouldnt have been deleted", v)
	}
	cache.Add("muhammad", []byte("is trying to become a deligent and amazing developer"))
	time.Sleep(waitTime)
	v, ok = cache.Get("muhammad")
	if ok {
		t.Errorf("%s should  have been deleted", v)
	}

}
