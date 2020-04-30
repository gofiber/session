// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package mysql

import (
	"fmt"
	"time"

	mysql "github.com/fasthttp/session/v2/providers/mysql"
)

// Config MySQL options
type Config struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Database  string
	TableName string

	Charset         string
	Collation       string
	Timeout         time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// New ...
func New(config ...Config) *mysql.Provider {
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
	if cfg.Charset == "" {
		cfg.Charset = "utf8"
	}
	if cfg.Collation == "" {
		cfg.Collation = "utf8_general_ci"
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 30 * time.Second
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 30 * time.Second
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 100
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 100
	}
	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = 1 * time.Second
	}

	provider, err := mysql.New(mysql.Config{
		Host:            cfg.Host,
		Port:            cfg.Port,
		Username:        cfg.Username,
		Password:        cfg.Password,
		Database:        cfg.Database,
		TableName:       cfg.TableName,
		Charset:         cfg.Charset,
		Collation:       cfg.Collation,
		Timeout:         cfg.Timeout,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		MaxIdleConns:    cfg.MaxIdleConns,
		MaxOpenConns:    cfg.MaxOpenConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	})
	if err != nil {
		fmt.Errorf("session: mysql %v", err)
	}
	return provider
}
