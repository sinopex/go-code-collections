// COPY FROM: https://go101.org/article/memory-leaking.html
package memory_leak

func WrongSubSlice(input []int, start, end int) []int {
	return input[start:end]
}

func GoodSubSlice(input []int, start, end int) []int {
	return append([]int{}, input[start:end]...)
}
