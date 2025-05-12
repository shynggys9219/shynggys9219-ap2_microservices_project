package safe

import (
	"context"
	"log"
	"runtime/debug"
)

func Do(ctx context.Context, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC RECOVERED] %v\n%s", r, debug.Stack())
		}
	}()

	fn()
}
