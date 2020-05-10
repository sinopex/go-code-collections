// 14. 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix/
package leetcode

func longestCommonPrefix(strs []string) string {
	var ptr int
	var wordCount = len(strs)
	if wordCount == 0 {
		return ""
	}
	running := true
	for running {
		// 遍历每一个单词
		for i := 0; i < wordCount; i++ {
			if ptr >= len(strs[i]) || (i > 0 && strs[i-1][ptr] != strs[i][ptr]) {
				running = false
				break
			}
		}
		if running {
			ptr++
		}
	}

	return strs[0][0:ptr]
}
