package example

import (
	"fmt"
	"testing"
	"time"
)

func TestUsageOfTimer(t *testing.T) {
	c := UsageOfTimer(func() { fmt.Println(time.Now()) }, time.Second)
	time.Sleep(time.Second * 2)
	close(c)
}
