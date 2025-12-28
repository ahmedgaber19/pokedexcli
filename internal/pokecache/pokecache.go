package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	CacheMap map[string]CacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		CacheMap: make(map[string]CacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheMap[key] = CacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.CacheMap[key]
	if !ok {
		return nil, false
	}
	return val.Val, true
}

func (c *Cache) reapLoop() {
	for range time.Tick(c.interval) {
		timeNow := time.Now()
		c.mu.Lock()
		for key, entry := range c.CacheMap {
			if timeNow.Sub(entry.CreatedAt) > c.interval {
				delete(c.CacheMap, key)
			}
		}
		c.mu.Unlock()

	}
}
