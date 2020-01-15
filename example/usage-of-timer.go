package example

import (
	"time"
)

func UsageOfTimer(fn func(), d time.Duration) chan struct{} {
	t := time.NewTimer(d)
	done := make(chan struct{})
	go func() {
		defer t.Stop()

		for {
			select {
			case <-t.C:
				fn()
				t.Reset(d)
			case <-done:
				return
			}
		}
	}()

	return done
}
