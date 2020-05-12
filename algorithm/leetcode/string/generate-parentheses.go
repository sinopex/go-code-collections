// 22 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/
package string

func generateParenthesis(n int) []string {
	var ret []string
	_generateParenthesis(n, n, "", &ret)
	return ret
}

func _generateParenthesis(left, right int, p string, ret *[]string) {
	if left == 0 && right == 0 {
		*ret = append(*ret, p)
		return
	}

	if left > 0 {
		_generateParenthesis(left-1, right, p+"(", ret)
	}

	if right > left && right > 0 {
		_generateParenthesis(left, right-1, p+")", ret)
	}
}
