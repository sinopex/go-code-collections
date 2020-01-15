// COPY FROM: https://go101.org/article/memory-leaking.html
package memory_leak

import (
	"strings"
)

func WrongSubString(input string, start, end int) string {
	return input[start:end]
}

func GoodSubString1(input string, start, end int) string {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(input[start:end])
	return b.String()
}

func GoodSubString2(input string, start, end int) string {
	return string([]byte(input[start:end]))
}
