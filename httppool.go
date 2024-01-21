package Wcache

import (
	"fmt"
	"github.com/wuerny/Wcache/consistenthash"
	"log/slog"
	"os"
	"sync"
)

//提供被其他节点通过http访问的能力

const (
	defaultBasePath = "/wcache/"
	defaultReplicas = 50
)

type HTTPPool struct {
	self        string
	basePath    string
	mu          sync.Mutex
	peers       *consistenthash.ConsistentHashMap
	httpGetters map[string]*httpGetter
}

func (p *HTTPPool) PickPeer(key string) (peer PeerGetter, ok bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	//log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
	baseLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	httppoolLogger := baseLogger.WithGroup("HTTPPool")
	httppoolLogger.Info("Info about HTTPPool", "Server", fmt.Sprintf(p.self), "info", fmt.Sprintf(format, v...))
}
