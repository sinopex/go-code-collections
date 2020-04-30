// 344 反转字符串
// https://leetcode-cn.com/problems/reverse-string/
package leetcode

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
