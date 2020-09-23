// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
// 🙏 Special thanks to @thomasvvugt & @savsgio (fasthttp/session)

package session

import (
	fsession "github.com/fasthttp/session/v2"
	"github.com/gofiber/fiber/v2"
)

// Store is a wrapper arround the session.Store
type Store struct {
	ctx  *fiber.Ctx
	sess *Session
	core *fsession.Store
}

// ID returns session id
func (s *Store) ID() string {
	return string(s.core.GetSessionID())
}

// Save storage before response
func (s *Store) Save() error {
	if err := s.sess.core.Save(s.ctx.Context(), s.core); err != nil {
		return err
	}
	if s.sess.config.noCookie {
		s.ctx.Response().Header.Del("Set-Cookie")
	}
	return nil
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

// Destroy session and cookies
func (s *Store) Destroy() error {
	return s.sess.core.Destroy(s.ctx.Context())
}

// Regenerate session id
func (s *Store) Regenerate() error {
	// https://github.com/fasthttp/session/blob/master/session.go#L205
	return s.sess.core.Regenerate(s.ctx.Context())
}
