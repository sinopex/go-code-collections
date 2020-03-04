package example

import (
	"fmt"
	"testing"
	"time"
)

func TestRunTicker(t *testing.T) {
	rt := RunTicker(func() { fmt.Println(time.Now()) }, time.Millisecond*1)
	time.Sleep(time.Millisecond * 2)
	// 如果chan关闭，立即会得到一个零值
	// 不用刻意传一个值，如：rt <- struct{}{}
	close(rt)
	time.Sleep(time.Millisecond * 10)
}
