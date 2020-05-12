// 226 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/
package tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l, r := root.Right, root.Left

	root.Left = invertTree(l)
	root.Right = invertTree(r)
	return root
}
