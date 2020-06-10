// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package postgres

import (
	"time"

	postgres "github.com/fasthttp/session/v2/providers/postgre"
	utils "github.com/gofiber/session/provider"
)

// Config Postgres options
type Config struct {
	Host            string
	Port            int64
	Username        string
	Password        string
	Database        string
	TableName       string
	Timeout         time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// New ...
func New(config ...Config) *postgres.Provider {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == 0 {
		cfg.Port = 3306
	}
	if cfg.Username == "" {
		cfg.Username = "root"
	}
	if cfg.Database == "" {
		cfg.Database = "session"
	}
	if cfg.TableName == "" {
		cfg.TableName = "session"
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
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
	provider, err := postgres.New(postgres.Config{
		Host:            cfg.Host,
		Port:            cfg.Port,
		Username:        cfg.Username,
		Password:        cfg.Password,
		Database:        cfg.Database,
		TableName:       cfg.TableName,
		Timeout:         cfg.Timeout,
		MaxIdleConns:    cfg.MaxIdleConns,
		MaxOpenConns:    cfg.MaxOpenConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	})
	if err != nil {
		utils.ErrorProvider("postgres", err)
	}
	return provider
}
