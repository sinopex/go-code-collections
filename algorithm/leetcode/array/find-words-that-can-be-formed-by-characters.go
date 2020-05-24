// 1160. 拼写单词
// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/
package array

func countCharacters(words []string, chars string) int {
	var m [26]int
	var result int
	for _, c := range chars {
		m[c-'a']++
	}

	for _, word := range words {
		t := m
		flag := true
		for _, v := range word {
			if t[v-'a'] == 0 {
				flag = false
				break
			}
			t[v-'a']--
		}

		if flag {
			result += len(word)
		}
	}

	return result
}
