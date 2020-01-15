package memory_leak

// ticker在stop时仅仅是关闭了计时器，并没有关闭ticker.C
// 而Ticker.C是只读类型，无法手动关闭，所以需要借助外力终止goroutine

import (
	"fmt"
	"log"
	"time"
)

func WrongTicker() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
			log.Println("tick")
		}
		log.Println("stopped")
	}()
	time.Sleep(3 * time.Second)
	log.Println("stopping ticker")
	ticker.Stop()
	time.Sleep(3 * time.Second)
}

func GoodTicker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("tick")
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()
	time.Sleep(3 * time.Second)
	log.Println("stopping ticker")
	done <- struct{}{}
	time.Sleep(3 * time.Second)
}
