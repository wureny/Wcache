package lru

import "container/list"

// TODO 采用LRU-k来缓解缓存污染问题
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

// 以下设计为了通用性
type Value interface {
	Len() int
}

// entry 是双向链表节点的数据类型
// 之所以要包含key，是为了方便删除字典中的映射
type entry struct {
	key   string
	value Value
}

func NewCache(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if e, ok := c.cache[key]; ok {
		// 将节点移动到队尾
		c.ll.MoveToFront(e)
		kv := e.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	// 获取队首节点
	e := c.ll.Back()
	if e != nil {
		// 删除节点
		c.ll.Remove(e)
		kv := e.Value.(*entry)
		// 删除map中的映射
		delete(c.cache, kv.key)
		// 更新当前占用内存
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		// 执行回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// TODO 以下函数的实现，使得数据大小有可能超过最大缓存限制
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
