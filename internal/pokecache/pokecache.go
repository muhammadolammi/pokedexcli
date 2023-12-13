package pokecache

import (
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]CacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMap: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheMap[key] = CacheEntry{
		CreatedAt: time.Now().UTC(),
		val:       val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.cacheMap[key]
	return cacheEntry.val, ok
}
func (c *Cache) reapLoop(interval time.Duration) {
	//create a ticker
	ticker := time.NewTicker(interval)
	//this send something new every interval so the for loop can fir
	for range ticker.C {
		c.reap(interval)
	}
}
func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().Add(-interval)
	for key, v := range c.cacheMap {
		if v.CreatedAt.Before(timeAgo) {
			delete(c.cacheMap, key)
		}
	}
}
