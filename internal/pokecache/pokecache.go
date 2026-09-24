package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mux          *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mux:          &sync.Mutex{},
	}

	go cache.reapLoop(duration)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cache := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cacheEntries[key] = cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	v, f := c.cacheEntries[key]
	if f {
		return v.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mux.Lock()
		for k, v := range c.cacheEntries {
			if time.Since(v.createdAt) > interval {
				delete(c.cacheEntries, k)
			}
		}
		c.mux.Unlock()
	}

}
