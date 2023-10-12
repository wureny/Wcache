package Wcache

// 代表缓存值的抽象与封装
// 由于是只读的，因此方法直接定义即可，不需要指针
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}
