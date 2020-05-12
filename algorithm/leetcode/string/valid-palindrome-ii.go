// 680. 验证回文字符串 Ⅱ
// https://leetcode-cn.com/problems/valid-palindrome-ii/
package string

func validPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}

	return _validPalindrome(s, true)
}
func _validPalindrome(s string, retry bool) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}

		if !retry {
			return false
		}

		return _validPalindrome(s[i+1:j+1], false) || _validPalindrome(s[i:j], false)
	}
	return true
}
