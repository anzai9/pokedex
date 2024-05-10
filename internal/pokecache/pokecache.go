package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
  mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
  cache := &Cache{
    cache: make(map[string]cacheEntry),
    mutex: &sync.Mutex{},
  }

  go cache.reapLoop(interval)

  return cache
}

func (c *Cache) Add(key string, val []byte) {
  c.mutex.Lock()
  defer c.mutex.Unlock()

  c.cache[key] = cacheEntry{
    createdAt: time.Now().UTC(),
    val: val,
  }
}

func (c *Cache) Get(key string) ([]byte, bool) {
  c.mutex.Lock()
  defer c.mutex.Unlock()

  val, exists := c.cache[key]
  if exists {
    return val.val, true
  }

  return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  defer ticker.Stop()

  for {
    select {
      case <- ticker.C:
        c.reap(interval)
    }
  }
}

func (c *Cache) reap(interval time.Duration) {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  for key, entry := range c.cache {
    if time.Since(entry.createdAt) > interval {
      delete(c.cache, key)
    }
  }
}
