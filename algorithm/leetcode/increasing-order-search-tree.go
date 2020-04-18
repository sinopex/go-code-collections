// 897 递增顺序查找树
// https://leetcode-cn.com/problems/increasing-order-search-tree/
package leetcode

func increasingBST(root *TreeNode) *TreeNode {
	var tree = &TreeNode{}
	var temp = tree

	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root.Left != nil {
			fn(root.Left)
		}

		tree := &TreeNode{
			Val:   root.Val,
			Right: nil,
			Left:  nil,
		}

		temp.Right = tree
		temp = temp.Right

		if root.Right != nil {
			fn(root.Right)
		}
	}

	fn(root)
	return tree.Right
}
