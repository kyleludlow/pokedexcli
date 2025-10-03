package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	data     map[string]CacheEntry
	interval time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(-c.interval)
		c.mu.Lock()
		for k, v := range c.data {
			if v.createdAt.Before(cutoff) {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}

var CACHE *Cache

func init() {
	CACHE = NewCache(30 * time.Second)
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]CacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}
