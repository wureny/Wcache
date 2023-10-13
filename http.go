package Wcache

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

//提供被其他节点通过http访问的能力

const defaultBasePath = "/wcache/"

type HTTPPool struct {
	self     string
	basePath string
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

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("unexpected url!")
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request!", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}
	val, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(val.ByteSlice())
}
