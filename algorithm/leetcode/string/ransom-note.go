// 383. 赎金信
// https://leetcode-cn.com/problems/ransom-note/
package string

func canConstruct(ransomNote string, magazine string) bool {
	var nodeLetters = [26]int{}
	for _, x := range []byte(magazine) {
		nodeLetters[x-'a']++
	}

	for _, x := range []byte(ransomNote) {
		nodeLetters[x-'a']--
		if nodeLetters[x-'a'] < 0 {
			return false
		}
	}

	return true
}
