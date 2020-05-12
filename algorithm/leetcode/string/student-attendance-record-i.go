// 551. 学生出勤记录 I
// https://leetcode-cn.com/problems/student-attendance-record-i/
package string

func checkRecord(s string) bool {
	var hasA bool
	var countL int
	for _, x := range s {
		if x == 'A' {
			if hasA {
				return false
			}
			hasA = true
			countL = 0
		} else if x == 'L' {
			countL++
			if countL == 3 {
				return false
			}
		} else {
			countL = 0
		}
	}
	return true
}
