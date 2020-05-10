// 819. 最常见的单词
// https://leetcode-cn.com/problems/most-common-word/
package leetcode

func mostCommonWord(paragraph string, banned []string) string {
	// 将禁用列表转换为hash，方便查找
	bannedMap := make(map[string]bool)
	for _, ban := range banned {
		bannedMap[ban] = true
	}
	words := make(map[string]int)
	n := len(paragraph)
	buf := []byte(paragraph)
	var isValidLetter = func(b byte) bool {
		return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
	}
	for i := 0; i < n; i++ {

		// 去除前导非字母字符
		for i < n && !isValidLetter(buf[i]) {
			i++
		}

		j := i
		// 按顺序找出单词
		for ; j < n; j++ {
			// 转成小写字母
			if buf[j] >= 'A' && buf[j] <= 'Z' {
				buf[j] += 32
			}
			// 遇到非字母字符，退出循环
			if !isValidLetter(buf[j]) {
				break
			}
		}

		if j > i {
			w := string(buf[i:j])
			if !bannedMap[w] {
				words[w]++
			}
		}

		i = j
	}

	max, ret := 0, ""
	for word, count := range words {
		if count > max {
			max = count
			ret = word
		}
	}
	return ret
}
