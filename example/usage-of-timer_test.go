package example

import (
	"testing"
	"time"
)

func TestUsageOfTimer(t *testing.T) {
	c := UsageOfTimer(func() {}, time.Millisecond*3)
	time.Sleep(time.Millisecond * 10)
	close(c)
}
