package Wcache

// 选择节点的接口
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 获取group中缓存值的接口
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
