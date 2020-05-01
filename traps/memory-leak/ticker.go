package memory_leak

// ticker在stop时仅仅是关闭了计时器，并没有关闭ticker.C
// 而Ticker.C是只读类型，无法手动关闭，所以需要借助外力终止goroutine

import (
	"time"
)

func WrongTicker() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
		}
	}()
	ticker.Stop()
}

func GoodTicker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
			case <-done:
				return
			}
		}
	}()
	done <- struct{}{}
}
