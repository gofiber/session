// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package memcache

import (
	"time"

	memcache "github.com/fasthttp/session/v2/providers/memcache"
	utils "github.com/gofiber/session/provider"
)

// Config memcache options
type Config struct {
	KeyPrefix    string
	ServerList   []string
	Timeout      time.Duration
	MaxIdleConns int
}

// New ...
func New(config ...Config) *memcache.Provider {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.KeyPrefix == "" {
		cfg.KeyPrefix = "session"
	}
	if len(cfg.ServerList) == 0 {
		cfg.ServerList = []string{
			"0.0.0.0:11211",
		}
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 8
	}
	provider, err := memcache.New(memcache.Config{
		KeyPrefix:    cfg.KeyPrefix,
		ServerList:   cfg.ServerList,
		Timeout:      cfg.Timeout,
		MaxIdleConns: cfg.MaxIdleConns,
	})
	if err != nil {
		utils.ErrorProvider("memcache", err)
	}
	return provider
}
