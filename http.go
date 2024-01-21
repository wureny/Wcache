package Wcache

import (
	"fmt"
	"log/slog"
	"os"
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
