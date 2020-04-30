// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package sqlite3

import (
	"fmt"
	"time"

	sqlite3 "github.com/fasthttp/session/v2/providers/sqlite3"
)

// Config redis options
type Config struct {
	DBPath          string
	TableName       string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// New ...
func New(config ...Config) *sqlite3.Provider {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.DBPath == "" {
		cfg.DBPath = "./"
	}
	if cfg.TableName == "" {
		cfg.TableName = "session"
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 100
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 100
	}
	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = 1 * time.Second
	}
	provider, err := sqlite3.New(sqlite3.Config{
		DBPath:          cfg.DBPath,
		TableName:       cfg.TableName,
		MaxIdleConns:    cfg.MaxIdleConns,
		MaxOpenConns:    cfg.MaxOpenConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	})
	if err != nil {
		fmt.Errorf("session: sqlite3 %v", err)
	}
	return provider
}
