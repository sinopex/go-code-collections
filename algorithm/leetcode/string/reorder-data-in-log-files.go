// 937. 重新排列日志文件
// https://leetcode-cn.com/problems/reorder-data-in-log-files/
package string

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	n := len(logs)
	// 数组切片，0存放标识符信息 1存放内容信息 2存放原文
	letterList := make([][3]string, 0, n)
	numberList := make([]string, 0, n)
	result := make([]string, 0, n)

	for _, log := range logs {
		pieces := strings.Fields(log)
		if pieces[1][0] >= '0' && pieces[1][0] <= '9' {
			numberList = append(numberList, log)
		} else {
			v := strings.Fields(log)
			letterList = append(letterList, [3]string{v[0], strings.Join(v[1:], " "), log})
		}
	}

	// 字母日志排序
	sort.Slice(letterList, func(i, j int) bool {
		// 内容相等，则根据标识符排序
		if letterList[i][1] == letterList[j][1] {
			return letterList[i][0] < letterList[j][0]
		}

		return letterList[i][1] < letterList[j][1]
	})

	// 压入字母日志
	for _, v := range letterList {
		result = append(result, v[2])
	}

	// 压入数字日志
	for _, n := range numberList {
		result = append(result, n)
	}
	return result
}
