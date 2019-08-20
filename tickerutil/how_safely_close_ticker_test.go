package tickerutil

import (
	"fmt"
	"testing"
	"time"
)

func TestRunTicker(t *testing.T) {
	rt := RunTicker(func() { fmt.Println(time.Now()) }, time.Second*1)
	time.Sleep(time.Second * 10)
	// 如果chan关闭，立即会得到一个零值
	// 不用刻意传一个值，如：rt <- struct{}{}
	close(rt)
	time.Sleep(time.Second)
}
