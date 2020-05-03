// 696. 计数二进制子串
// https://leetcode-cn.com/problems/count-binary-substrings/
package leetcode

func countBinarySubstrings(s string) int {
	var startByte byte
	var startNum, matchNum, cnt int
	var flag int

	for i := 0; i < len(s); {
		startByte, startNum, matchNum, flag = s[i], 1, 0, 1
		for j := i + 1; j < len(s); {
			if s[j] == startByte {
				if flag == 1 {
					startNum++
					j++
				} else {
					break
				}
			} else {
				flag = 2
				matchNum++
				j++
			}
		}
		if startNum > matchNum {
			cnt += matchNum
		} else {
			cnt += startNum
		}
		i += startNum

		if flag == 1 {
			break
		}
	}

	return cnt
}
