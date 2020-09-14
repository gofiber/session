// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package session

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	fsession "github.com/fasthttp/session/v2"
	"github.com/fasthttp/session/v2/providers/memory"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// Config defines the config for RequestID middleware
type Config struct {
	// Lookup is a string in the form of "<source>:<name>" that is used
	// to extract session id from the request.
	// Possible values: "header:<name>", "query:<name>" or "cookie:<name>"
	// Optional. Default value "cookie:session_id".
	Lookup string

	// Secure attribute for cookie
	// Optional. Default: false
	Secure bool

	// Domain attribute for cookie
	// Optional. Default: ""
	Domain string

	// SameSite attribute for cookie
	// Possible values: "Lax", "Strict", "None"
	// Optional. Default: "Lax"
	SameSite string

	//  0 means no expire, (24 years)
	// -1 means when browser closes
	// >0 is the time.Duration which the session cookies should expire.
	// Optional. Default: 12 hours
	Expiration time.Duration

	// Holds the provider interface
	// Optional. Default: memory.New()
	Provider fsession.Provider

	// Generator is a function that generates an unique id
	// Optional.
	Generator func() []byte

	// gc life time to execute it
	// Optional. 1 minute
	GCInterval time.Duration

	// fasthttp/session always sets a cookie response
	// we want to disable this if no cookie lookup is set
	noCookie bool
}

// Session ...
type Session struct {
	config    Config
	core      *fsession.Session
	storePool *sync.Pool
}

// New ...
func New(config ...Config) *Session {
	// Init session config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	// Fiber Config
	if cfg.Lookup == "" {
		cfg.Lookup = "cookie:session_id"
	}
	if cfg.Expiration == 0 {
		cfg.Expiration = 12 * time.Hour
	}
	if cfg.GCInterval == 0 {
		cfg.GCInterval = 1 * time.Minute
	}
	if cfg.Generator == nil {
		cfg.Generator = defaultGenerator
	}
	if cfg.Provider == nil {
		provider, err := memory.New(memory.Config{})
		if err != nil {
			log.Fatalf("session: memory %v", err)
		}
		cfg.Provider = provider
	}

	// Private fasthttp config
	var scfg fsession.Config
	scfg.GCLifetime = cfg.GCInterval
	scfg.SessionIDGeneratorFunc = cfg.Generator
	// Split lookup key <method>:<key>
	parts := strings.Split(cfg.Lookup, ":")
	// Cookie configuration for fasthttp
	scfg.CookieName = parts[1]
	scfg.Domain = cfg.Domain
	scfg.Expiration = cfg.Expiration
	scfg.Secure = cfg.Secure
	switch strings.ToLower(cfg.SameSite) {
	case "strict":
		scfg.CookieSameSite = fasthttp.CookieSameSiteStrictMode
	case "none":
		scfg.CookieSameSite = fasthttp.CookieSameSiteNoneMode
	default:
		scfg.CookieSameSite = fasthttp.CookieSameSiteLaxMode
	}
	// Set configuration for header and query lookups
	scfg.SessionIDInHTTPHeader = parts[0] == "header"
	scfg.SessionNameInHTTPHeader = parts[1]
	scfg.SessionIDInURLQuery = parts[0] == "query"
	scfg.SessionNameInURLQuery = parts[1]
	// fasthttp/session always sets a cookie response
	// we want to disable this if no cookie lookup is set
	if parts[0] != "cookie" {
		cfg.noCookie = true
	}
	// Create fiber session
	sessions := &Session{
		config: cfg,
		core:   fsession.New(scfg),
	}
	provider := fmt.Sprintf("%v", reflect.TypeOf(cfg.Provider))
	switch provider {
	case "*mysql.Provider", "*sqlite3.Provider":
		scfg.EncodeFunc = fsession.Base64Encode
		scfg.DecodeFunc = fsession.Base64Decode
	default:
		// redis / memcache / memory
		scfg.EncodeFunc = fsession.MSGPEncode
		scfg.DecodeFunc = fsession.MSGPDecode
	}

	// Set default provider
	if err := sessions.core.SetProvider(cfg.Provider); err != nil {
		log.Fatal(err)
	}

	return sessions
}

// Get store
func (s *Session) Get(ctx *fiber.Ctx) *Store {
	fstore, _ := s.core.Get(ctx.Context())
	return &Store{
		ctx:  ctx,
		sess: s,
		core: fstore,
	}
}

func defaultGenerator() []byte {
	return utils.GetBytes(utils.UUID())
}
