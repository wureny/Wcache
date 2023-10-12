package Wcache

import (
	"github.com/wuerny/Wcache/lru"
	"sync"
)

//添加了并发特性

type Cache struct {
	mu         *sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func NewCache(maxBytes int64, m *sync.Mutex) *Cache {
	return &Cache{
		lru:        lru.NewCache(maxBytes, nil),
		mu:         m,
		cacheBytes: maxBytes,
	}
}

func (c *Cache) Add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lru.Add(key, value)
}

func (c *Cache) Get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
