// 938 二叉搜索树的范围和
// https://leetcode-cn.com/problems/range-sum-of-bst/
package tree

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val >= L && root.Val <= R {
		res += root.Val
	}

	// 合理运用二叉搜索树的特性，进行剪枝
	// 如果节点值比目标左值小，则只用搜索右边
	if root.Val >= L {
		res += rangeSumBST(root.Left, L, R)
	}

	// 同理，如果节点值比目标右值大，则只用搜索左边
	if root.Val <= R {
		res += rangeSumBST(root.Right, L, R)
	}

	return res
}
