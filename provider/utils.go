package provider

import (
	"fmt"
	"log"
)

// ErrorProvider ...
func ErrorProvider(provider string, err error) {
	if e := fmt.Errorf("session: %v %v", provider, err); e != nil {
		log.Printf("session error: %v; provider: %v;", e, provider)
	}
}
