// 125. 验证回文串
// https://leetcode-cn.com/problems/valid-palindrome/
package string

func isPalindrome(s string) bool {
	var isValidLetter = func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}

	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		for !isValidLetter(s[i]) && i < j {
			i++
		}
		for !isValidLetter(s[j]) && j > i {
			j--
		}

		// 左右都转换成小写字母
		// fmt.Printf("%d=%s,%d=%s\n", i, string(s[i]), j, string(s[j]))
		l, r := s[i], s[j]
		if l >= 'A' && l <= 'Z' {
			l += 32
		}
		if r >= 'A' && r <= 'Z' {
			r += 32
		}

		if l != r {
			return false
		}
	}

	return true
}
