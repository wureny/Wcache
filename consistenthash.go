package Wcache

import (
	"hash/crc32"
	"sort"
)

// 一致性哈希算法
type Hash func(data []byte) uint32

// 一致性哈希算法的实现
type Map struct {
	hash Hash
	// 虚拟节点倍数
	replicas int
	keys     []int // Sorted
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	return &Map{
		hash:     fn,
		replicas: replicas,
		keys:     nil,
		hashMap:  nil,
	}
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		// 每个真实节点创建m.replicas个虚拟节点
		for i := 0; i < m.replicas; i++ {
			// 虚拟节点的名称为strconv.Itoa(i) + key
			hash := int(m.hash([]byte(string(i) + key)))
			// 添加到环上
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	// 环上的虚拟节点排序
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	// 计算key的哈希值
	hash := int(m.hash([]byte(key)))
	// 顺时针找到第一个匹配的虚拟节点的下标idx
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 通过hashMap映射得到真实节点的名称
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
