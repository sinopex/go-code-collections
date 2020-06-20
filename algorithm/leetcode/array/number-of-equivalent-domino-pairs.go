// 1128. 等价多米诺骨牌对的数量
// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/
package array

import (
	"sort"
)

func numEquivDominoPairs(dominoes [][]int) int {
	count := 0
	arr := make([]int, 101)
	for i := 0; i < len(dominoes); i++ {
		a, b := dominoes[i][0], dominoes[i][1]
		if a > b {
			a, b = b, a
		}
		count += arr[a*10+b]
		arr[a*10+b]++
	}

	return count
}
func numEquivDominoPairs1(dominoes [][]int) int {
	if len(dominoes) <= 1 {
		return 0
	}
	s := make([]int, len(dominoes))
	for k, v := range dominoes {
		if v[1] > v[0] {
			s[k] = v[1]*10 + v[0]
		} else {
			s[k] = v[0]*10 + v[1]
		}
	}

	sort.Ints(s)
	count := 0
	prev := s[0]
	for i := 1; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[j] == prev {
				count++
			} else {
				break
			}
		}
		prev = s[i]
	}

	return count
}
