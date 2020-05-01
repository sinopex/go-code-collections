// 1189. “气球” 的最大数量
// https://leetcode-cn.com/problems/maximum-number-of-balloons/
package leetcode

import (
	"math"
)

func maxNumberOfBalloons(text string) int {
	var dict = map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}

	for _, r := range text {
		if _, ok := dict[r]; ok {
			dict[r]++
		}
	}

	var n = math.MaxInt32

	for r, v := range dict {
		if r == 'l' || r == 'o' {
			v /= 2
		}
		if v < n {
			n = v
		}
	}

	return n
}
