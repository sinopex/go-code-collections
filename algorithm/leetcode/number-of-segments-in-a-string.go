// 434. 字符串中的单词数
// https://leetcode-cn.com/problems/number-of-segments-in-a-string/
package leetcode

func countSegments(s string) int {
	count, i, n := 0, 0, len(s)
	// 删除字符串的前置空格
	for i < n && s[i] == ' ' {
		i++
	}
	for ; i < n; i++ {
		// 排完一个单词
		hasWorld := false
		for i < n {
			// 遇到空格停止
			if s[i] == ' ' {
				break
			}
			hasWorld = true
			i++
		}
		if hasWorld {
			count++
		}
	}

	return count
}
