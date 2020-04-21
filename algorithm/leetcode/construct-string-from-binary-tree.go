// 606 根据二叉树创建字符串
// https://leetcode-cn.com/problems/construct-string-from-binary-tree/
package leetcode

import (
	"strconv"
)

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")(" + tree2str(t.Right) + ")"
}
