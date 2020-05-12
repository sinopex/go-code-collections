// 1022 从根到叶的二进制之和
// https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/
package tree

func sumRootToLeaf(root *TreeNode) int {
	// 如果这里 没有重新 给这个全局变量 赋值，下一次运行时，这个全局变量保留上次运行的值
	var sumRootToLeafRes int
	var fn func(root *TreeNode, n int)
	fn = func(root *TreeNode, n int) {
		n <<= 1
		n += root.Val
		if root.Left == nil && root.Right == nil {
			sumRootToLeafRes += n
			return
		}

		if root.Left != nil {
			fn(root.Left, n)
		}
		if root.Right != nil {
			fn(root.Right, n)
		}
	}

	fn(root, 0)
	return sumRootToLeafRes
}
