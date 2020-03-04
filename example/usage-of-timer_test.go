package example

import (
	"fmt"
	"testing"
	"time"
)

func TestUsageOfTimer(t *testing.T) {
	c := UsageOfTimer(func() { fmt.Println(time.Now()) }, time.Millisecond*3)
	time.Sleep(time.Millisecond * 10)
	close(c)
}
