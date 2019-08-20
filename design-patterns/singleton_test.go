package design_patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()
	s1.Set(100)
	assert.Equal(t, GetInstance().Get(), 100)
	s2.Set(200)
	assert.NotEqual(t, GetInstance().Get(), 300)
}
