package array

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	rows := make(map[int][]int, len(mat))
	arr := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		c := 0
		for _, v := range mat[i] {
			if v&1 == 0 {
				break
			}
			c++
		}
		if rows[c] == nil {
			rows[c] = []int{}
		}
		rows[c] = append(rows[c], i)
		arr[i] = c
	}
	sort.Ints(arr)
	result := make([]int, k)
	c := 0
	prev := -1
	for _, v := range arr {
		if v == prev {
			continue
		}
		prev = v
		for _, vv := range rows[v] {
			result[c] = vv
			if c == k-1 {
				goto END
			}
			c++
		}
	}
END:
	return result
}
