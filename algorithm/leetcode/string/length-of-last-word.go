// 58. 最后一个单词的长度
// https://leetcode-cn.com/problems/length-of-last-word/
package string

func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	j := n - 1
	// 去掉末位的空格
	for j >= 0 && s[j] == ' ' {
		j--
	}

	count := 0
	for j >= 0 {
		if s[j] == ' ' {
			return count
		}
		count++
		j--
	}

	return count
}
