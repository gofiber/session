package redis

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/fasthttp/session/v2/providers/redis"
)

// Config Redis options
type Config struct {
	KeyPrefix          string
	Network            string
	Addr               string
	Password           string
	DB                 int
	MaxRetries         int
	MinRetryBackoff    time.Duration
	MaxRetryBackoff    time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
	TLSConfig          *tls.Config
	// Limiter            redis.Limiter
}

// New ...
func New(config ...Config) *redis.Provider {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.KeyPrefix == "" {
		cfg.KeyPrefix = "session"
	}
	if cfg.Addr == "" {
		cfg.Addr = "127.0.0.1:6379"
	}
	if cfg.PoolSize == 0 {
		cfg.PoolSize = 8
	}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = 30 * time.Second
	}

	provider, err := redis.New(redis.Config{
		KeyPrefix:          cfg.KeyPrefix,
		Network:            cfg.Network,
		Addr:               cfg.Addr,
		Password:           cfg.Password,
		DB:                 cfg.DB,
		MaxRetries:         cfg.MaxRetries,
		MinRetryBackoff:    cfg.MinRetryBackoff,
		MaxRetryBackoff:    cfg.MaxRetryBackoff,
		DialTimeout:        cfg.DialTimeout,
		ReadTimeout:        cfg.ReadTimeout,
		WriteTimeout:       cfg.WriteTimeout,
		PoolSize:           cfg.PoolSize,
		MinIdleConns:       cfg.MinIdleConns,
		MaxConnAge:         cfg.MaxConnAge,
		PoolTimeout:        cfg.PoolTimeout,
		IdleTimeout:        cfg.IdleTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
		TLSConfig:          cfg.TLSConfig,
		// Limiter             cfg.Limiter,
	})
	if err != nil {
		fmt.Errorf("session: redis %v", err)
	}
	return provider
}
