package cacheGo

import (
	"fmt"
	"sync"
	"time"
)

type cacheItem struct {
	value int
	ttl   time.Time
}

type Cache struct {
	items map[string]cacheItem
	mu    sync.Mutex
}

func New() *Cache {
	return &Cache{items: make(map[string]cacheItem)}
}

func (c *Cache) Get(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]

	if !found || time.Now().After(item.ttl) {
		return 0, false
	}

	return item.value, true
}

func (c *Cache) Set(key string, value int, ttl time.Duration) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.items == nil {
		c.items = make(map[string]cacheItem)
	}

	c.items[key] = cacheItem{
		value: value,
		ttl:   time.Now().Add(ttl),
	}

	return fmt.Sprintf("Added %v with %v, expires in %v\n", key, value, ttl)
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *Cache) CleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, item := range c.items {
		if time.Now().After(item.ttl) {
			delete(c.items, key)
		}
	}
}
