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

// l可以为nil，表示使用延迟初始化的方案
func NewCache(l *lru.Cache, maxBytes int64) *Cache {
	return &Cache{
		lru:        l,
		cacheBytes: maxBytes,
	}
}

func (c *Cache) Add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//延迟初始化
	if c.lru == nil {
		c.lru = lru.NewCache(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *Cache) Get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//延迟初始化
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
