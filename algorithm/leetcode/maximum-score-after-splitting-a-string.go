// 1422. 分割字符串的最大得分
// https://leetcode-cn.com/problems/maximum-score-after-splitting-a-string/
package leetcode

func maxScore(s string) int {
	var max, zero, one int
	buf := []byte(s)
	for _, v := range buf {
		if v == '1' {
			one++
		}
	}

	for i := 0; i < len(buf)-1; i++ {
		if buf[i] == '0' {
			zero++
		} else {
			one--
		}

		if zero+one > max {
			max = zero + one
		}
	}
	return max
}
