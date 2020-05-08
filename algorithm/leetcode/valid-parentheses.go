// 20. 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/
package leetcode

func isValid(s string) bool {
	var stack = make([]byte, 0)
	var n int
	var dict = map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	for _, r := range []byte(s) {
		// 如果是开始标签，则进栈
		if _, ok := dict[r]; ok {
			stack = append(stack, r)
			n++
			continue
		}
		// 否则为闭合标签
		// case 1：当前栈为空
		if n == 0 {
			return false
		}

		// case 2: 当前栈顶元素与闭合标签无法配对
		if dict[stack[n-1]] != r {
			return false
		}
		// case 3: 匹配成功，栈长度减1
		stack = stack[0 : n-1]
		n--
	}

	return n == 0
}
