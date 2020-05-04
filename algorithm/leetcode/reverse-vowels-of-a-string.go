// 345. 反转字符串中的元音字母
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
package leetcode

func reverseVowels(s string) string {
	var buf = []byte(s)
	var dict = map[byte]bool{
		'a': true, 'o': true, 'e': true, 'i': true, 'u': true,
		'A': true, 'O': true, 'E': true, 'I': true, 'U': true,
	}
	for l, r, findL, findR := 0, len(buf)-1, false, false; l < r; {
		if dict[buf[l]] {
			findL = true
		} else {
			l++
			findL = false
		}

		if dict[buf[r]] {
			findR = true
		} else {
			r--
			findR = false
		}

		if findL && findR {
			buf[l], buf[r] = buf[r], buf[l]
			l++
			r--
		}
	}

	return string(buf)
}
