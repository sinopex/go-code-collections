// 538 把二叉搜索树转换为累加树
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
package tree

func convertBST(root *TreeNode) *TreeNode {
	var n int
	_convertBST(root, &n)
	return root
}
func _convertBST(root *TreeNode, n *int) {
	if root == nil {
		return
	}

	_convertBST(root.Right, n)
	root.Val += *n
	*n = root.Val
	_convertBST(root.Left, n)
}
