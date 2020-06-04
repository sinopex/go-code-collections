// 1185. 一周中的第几天
// https://leetcode-cn.com/problems/day-of-the-week/
package array

import (
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	var y = 1971
	var d = 0
	var res int
	str := [7]string{"Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	monthDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for ; y < year; y++ {
		if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
			res += 366
		} else {
			res += 365
		}
	}
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		monthDays[1] = 29
	}
	for ; d < month-1; d++ {
		res += monthDays[d]
	}
	res += day
	res = (res - 1) % 7
	return str[res]
}

func dayOfTheWeek1(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Weekday().String()
}
