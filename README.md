# Session

![Test](https://github.com/gofiber/session/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/session/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/session/workflows/Linter/badge.svg)

This session middleware is build on top of [fasthttp/session](https://github.com/fasthttp/session) by [@savsgio](https://github.com/savsgio) [MIT](https://github.com/fasthttp/session/blob/master/LICENSE).
Special thanks to [@thomasvvugt](https://github.com/thomasvvugt) for helping with this middleware.

### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/session
```

### Signature
```go
session.New(config ...session.Config) *Session
```

### Config
| Property | Type | Description | Default |
| :--- | :--- | :--- | :--- |
| Lookup | `string` | Where to look for the session id `<source>:<name>`, possible values: `cookie:key`, `header:key`, `query:key` | `"cookie:session_id"` |
| Domain | `string` | Cookie domain | `""` |
| Expiration | `time.Duration` | Session expiration time, possible values : `0` means no expiry (24 years), `-1` means when the browser closes, `>0` is the `time.Duration` which the session cookies should expire. | `12 * time.Hour` |
| Secure | `bool` | If the cookie should only be send over HTTPS | `false` |
| Provider | `Provider` | Holds the provider interface | `memory.Provider` |
| Generator | `func() []byte` | Generator is a function that generates an unique id | `uuid` |
| GCInterval | `time.Duration` | Interval for the garbage collector | `uuid` |

### Default Example
```go
package main

import (
  "fmt"

  "github.com/gofiber/fiber"
  "github.com/gofiber/session"
)

func main() {
  app := fiber.New()

  // create session handler
  sessions := session.New()

  app.Get("/", func(c *fiber.Ctx) {
    store := sessions.Get(c)    // get/create new session
    defer store.Save()

    store.ID()                   // returns session id
    store.Destroy()              // delete storage + cookie
    store.Get("john")            // get from storage
    store.Regenerate()           // generate new session id
    store.Delete("john")         // delete from storage
    store.Set("john", "doe")     // save to storage
  })
  
  app.Listen(3000)
}
```

### Provider Example
```go
package main

import (
  "fmt"

  "github.com/gofiber/fiber"
  "github.com/gofiber/session"
  "github.com/gofiber/session/provider/memcache"
  // "github.com/gofiber/session/provider/mysql"
  // "github.com/gofiber/session/provider/postgres"
  // "github.com/gofiber/session/provider/redis"
  // "github.com/gofiber/session/provider/sqlite3"
)

func main() {
  app := fiber.New()

  provider := memcache.New(memcache.Config{
    KeyPrefix:  "session",
    ServerList: []string{
      "0.0.0.0:11211",
    },
    MaxIdleConns: 8,
  })
  // provider := mysql.New(mysql.Config{
  //   Host:       "session",
  //   Port:       3306,
  //   Username:   "root",
  //   Password:   "",
  //   Database:   "test",
  //   TableName:  "session",
  // })
  // provider := postgres.New(postgres.Config{
  //   Host:       "session",
  //   Port:       5432,
  //   Username:   "root",
  //   Password:   "",
  //   Database:   "test",
  //   TableName:  "session",
  // })
  // provider := redis.New(redis.Config{
  //   KeyPrefix:   "session",
  //   Addr:        "127.0.0.1:6379",
  //   PoolSize:    8,
  //   IdleTimeout: 30 * time.Second,
  // })
  // provider := sqlite3.New(sqlite3.Config{
  //   DBPath:     "test.db",
  //   TableName:  "session",
  // })

  sessions := session.New(session.Config{
    Provider: provider,
  })

  app.Get("/", func(c *fiber.Ctx) {
    store := sessions.Get(c)    // get/create new session
    defer store.Save()

    store.ID()                   // returns session id
    store.Destroy()              // delete storage + cookie
    store.Get("john")            // get from storage
    store.Regenerate()           // generate new session id
    store.Delete("john")         // delete from storage
    store.Set("john", "doe")     // save to storage
  })
  
  app.Listen(3000)
}
```
