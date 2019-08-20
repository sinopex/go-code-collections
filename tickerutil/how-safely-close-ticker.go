package tickerutil

import (
	"time"
)

// 如何正确的关闭定时器
func RunTicker(fn func(), d time.Duration) chan struct{} {
	t := time.NewTicker(d)
	done := make(chan struct{})

	go func() {
		defer t.Stop()

		for {
			select {
			case <-t.C:
				fn()
			case <-done:
				return
			}
		}
	}()

	return done
}
