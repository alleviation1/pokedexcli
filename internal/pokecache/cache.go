package pokecache

import(
	"time"
	"sync"
)

type Cache struct {
	// api call as string key : cache Entry containing time and data
	cacheEntries	map[string]cacheEntry
	mu				*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, res []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now().UTC(),
		val: res,
	}
	c.cacheEntries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cacheEntries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheEntries, k)
		}
	}
}