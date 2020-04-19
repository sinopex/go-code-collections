// 257 二叉树的所有路径
// https://leetcode-cn.com/problems/binary-tree-paths/
package leetcode

import (
	"fmt"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	ret := make([]string, 0)
	_binaryTreePaths(root, &ret, "", "")
	return ret
}
func _binaryTreePaths(root *TreeNode, ret *[]string, path, prefix string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, fmt.Sprintf("%s%s%d", path, prefix, root.Val))
		return
	}

	_binaryTreePaths(root.Left, ret, fmt.Sprintf("%s%s%d", path, prefix, root.Val), "->")
	_binaryTreePaths(root.Right, ret, fmt.Sprintf("%s%s%d", path, prefix, root.Val), "->")
}
