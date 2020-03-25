// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package session

import (
	"time"

	"github.com/fasthttp/session"
	"github.com/fasthttp/session/memory"
	"github.com/gofiber/fiber"
)

// Config defines the config for RequestID middleware
type Config struct {
	// Session key
	// Optional. Default: "sessionid"
	Key string
	// Where do store the session ID?
	// Possible values: "cookie", "header", "query"
	// Optional. Default: "cookie".
	Lookup string
	// Cookie domain
	// Optional. Default: ""
	Domain string
	// Expire time
	// Optional. Default: 2 * time.Hour
	Expires time.Duration
	// Secure cookie
	// Optional. Default: false
	Secure bool
	// Session ID generator
	// Optional. Default nil
	Generator func() []byte
}

// Session is a wrapper arround the session.Session
type Session struct {
	core *session.Session
}

// Store is a wrapper arround the session.Storer
type Store struct {
	ctx  *fiber.Ctx
	sess *Session
	core session.Storer
}

// New adds an indentifier to the request using the `X-Request-ID` header
func New(config ...Config) *Session {
	// Init session config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Key == "" {
		cfg.Key = "sessionid"
	}
	if cfg.Lookup == "" {
		cfg.Lookup = "cookie"
	}
	if cfg.Expires == 0 {
		cfg.Expires = 2 * time.Hour
	}

	// Fasthttpp session config
	var scfg session.Config

	scfg.CookieName = cfg.Key
	scfg.Domain = cfg.Domain
	scfg.Expires = cfg.Expires
	scfg.Secure = cfg.Secure

	scfg.SessionIDGeneratorFunc = cfg.Generator
	scfg.SessionIDInHTTPHeader = cfg.Lookup == "header"
	scfg.SessionNameInHTTPHeader = cfg.Key
	scfg.SessionIDInURLQuery = cfg.Lookup == "query"
	scfg.SessionNameInURLQuery = cfg.Key

	sessions := &Session{
		core: session.New(&scfg),
	}

	// Set default provider
	sessions.core.SetProvider("memory", &memory.Config{})

	return sessions
}

// Start get/sets session storage from Ctx
func (s *Session) Start(ctx *fiber.Ctx) *Store {
	fstore, _ := s.core.Get(ctx.Fasthttp)
	return &Store{
		ctx:  ctx,
		sess: s,
		core: fstore,
	}
}

// Save storage before response
func (s *Store) Save(c *fiber.Ctx, store *Store) {
	s.sess.core.Save(c.Fasthttp, store.core)
}

// Get get data by key
func (s *Store) Get(key string) interface{} {
	return s.core.Get(key)
}

// Set get data by key
func (s *Store) Set(key string, value interface{}) {
	s.core.Set(key, value)
}

// Delete delete data by key
func (s *Store) Delete(key string) {
	s.core.Delete(key)
}

// Empty storage
func (s *Store) Empty() {
	s.core.Flush()
}

// Destroy session
func (s *Store) Destroy() {
	s.sess.core.Destroy(s.ctx.Fasthttp)
}

// Regenerate session id
func (s *Store) Regenerate() {
	fstore, _ := s.sess.core.Regenerate(s.ctx.Fasthttp)
	s.core = fstore
}

// Expires set expiration for the session
func (s *Store) Expires(expiration ...time.Duration) (t time.Duration) {
	if len(expiration) > 0 {
		s.core.SetExpiration(expiration[0])
		return
	}
	return s.core.GetExpiration()
}

// ID returns session id
func (s *Store) ID() string {
	return string(s.core.GetSessionID())
}
