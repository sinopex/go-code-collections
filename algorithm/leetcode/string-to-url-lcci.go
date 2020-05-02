// 面试题 01.03. URL化
// https://leetcode-cn.com/problems/string-to-url-lcci/
package leetcode

func replaceSpaces(S string, length int) string {
	var s = make([]byte, 0, len(S))
	for _, v := range []byte(S) {
		if v == ' ' {
			if length > 0 {
				s = append(s, '%', '2', '0')
			}
		} else {
			s = append(s, v)
		}
		length--
	}

	return string(s)
}
