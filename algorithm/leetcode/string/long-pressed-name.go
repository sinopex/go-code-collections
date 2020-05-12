// 925. 长按键入
// https://leetcode-cn.com/problems/long-pressed-name/
package string

func isLongPressedName(name string, typed string) bool {
	n1 := len(name)
	n2 := len(typed)
	i, j := 0, 0
	for i < n1 && j < n2 {
		// 如果值不一样，返回错误
		if name[i] != typed[j] {
			return false
		}
		countI, countJ := 0, 0

		// 过掉后面值相同的多余字段
		for i = i + 1; i < n1 && name[i] == name[i-1]; {
			i++
			countI++
		}

		// 过掉后面值相同的多余字段
		for j = j + 1; j < n2 && typed[j] == typed[j-1]; {
			countJ++
			j++
		}

		// 如果原始重复字段比长按的字段个数要多，返回错误
		if countI > countJ {
			return false
		}
	}

	return i == n1 && j == n2
}
