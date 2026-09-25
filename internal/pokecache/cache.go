package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	mu       sync.RWMutex
	cache    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{cache: make(map[string]cacheEntry), interval: interval}
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			reapLoop(cache)
		}
	}()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{val: val, createAt: time.Now()}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return
	}

	val = entry.val
	return
}

func reapLoop(c *Cache) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for k, v := range c.cache {
		if now.Sub(v.createAt) > c.interval {
			delete(c.cache, k)
		}
	}
}
