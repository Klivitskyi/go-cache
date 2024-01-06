package cacheGo

import "fmt"

type Cache struct {
	make map[string]int
}

func New() Cache {
	return Cache{make: make(map[string]int)}
}

func (c *Cache) Get(key string) int {
	return c.make[key]
}

func (c *Cache) Set(key string, value int) string {
	if c.make == nil {
		c.make = make(map[string]int)
	}

	c.make[key] = value
	return fmt.Sprintf("Added %v with %v", key, value)
}

func (c *Cache) Delete(key string) {
	delete(c.make, key)
}
