// 938 二叉搜索树的范围和
// https://leetcode-cn.com/problems/range-sum-of-bst/
package leetcode

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val >= L && root.Val <= R {
		res += root.Val
	}

	return res + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
