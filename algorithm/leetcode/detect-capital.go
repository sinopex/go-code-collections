// 520. 检测大写字母
// https://leetcode-cn.com/problems/detect-capital/
package leetcode

func detectCapitalUse(word string) bool {
	var mode = 3 // 只有三种可能性：即：1全小写 2全大写 3 首字母大写
	for i, x := range word {
		if i == 0 {
			// 进入模式3
			if x < 'A' || x > 'Z' {
				mode = 1
				continue
			}
		} else if i == 1 {
			// 第二个字母遇到大写
			if x >= 'A' && x <= 'Z' {
				if mode == 1 {
					// 如果是模式1，则返回失败
					return false
				} else {
					// 转换为模式2
					mode = 2
				}
			}
		} else {
			// 模式2后续不能有小写
			if mode == 2 && (x < 'A' || x > 'Z') {
				return false
			}
			// 模式1和3，后续不能有大写
			if (mode == 3 || mode == 1) && (x >= 'A' && x <= 'Z') {
				return false
			}
		}
	}
	return true
}
