### Session
...

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
| Key | `string` | Defines a function to skip middleware | `"sessionid"` |
| Lookup | `string` | Where to look for the session id, possible values: `cookie`, `header`, `query` | `"cookie"` |
| Domain | `string` | Cookie domain | `""` |
| Expires | `time.Duration` | Session expiry | `2 * time.Hour` |
| Secure | `bool` | Custom response body for unauthorized responses | `false` |

### Example
```go
package main

import (
  "fmt"

  "github.com/gofiber/fiber"
  "github.com/gofiber/session"
)

func main() {
  app := fiber.New()
  
  // optional config
  config := session.Config{
    Key:    "dinosaurus",       // default: "sessionid"
    Lookup: "header",           // default: "cookie"
    Domain: "google.com",       // default: ""
    Expires: 30 * time.Minutes, // default: 2 * time.Hour
    Secure:  true,              // default: false
  }

  // create session handler
  sessions := session.New(config)

  app.Get("/", func(c *fiber.Ctx) {
    store := sessions.Start(c)   // get/create new session

    store.ID()                   // returns session id
    store.Empty()                // empty storage
    store.Destroy()              // delete storage + cookie
    store.Get("john")            // get from storage
    store.Regenerate()           // generate new session id
    store.Delete("john")         // delete from storage
    store.Set("john", "doe")     // save to storage
    store.Expires(2 * time.Hour) // set session expiration
  })
  
  app.Listen(3000)
}
```