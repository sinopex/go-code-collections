// 1408. 数组中的字符串匹配
// https://leetcode-cn.com/problems/string-matching-in-an-array/
package string

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	n := len(words)
	result := make([]string, 0, n-1)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}
