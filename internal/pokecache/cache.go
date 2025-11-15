package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}
type PokeCache struct {
	Cache    map[string]CacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *PokeCache {
	c := &PokeCache{
		Cache:    map[string]CacheEntry{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *PokeCache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = CacheEntry{Val: val, CreatedAt: time.Now().UTC()}
}

func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.Cache[key]
	if !ok {
		return nil, false
	}
	return val.Val, true
}

func (c *PokeCache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		timeNow := time.Now().UTC()
		for key, val := range c.Cache {
			if timeNow.Sub(val.CreatedAt) > c.interval {
				delete(c.Cache, key)
			}
		}
		c.mu.Unlock()

	}

}
