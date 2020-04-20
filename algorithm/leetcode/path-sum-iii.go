// 437 路径总和
// https://leetcode-cn.com/problems/path-sum-iii/
package leetcode

func PathFrom(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0

	if root.Val == sum {
		cnt++
	}

	cnt += PathFrom(root.Left, sum-root.Val)
	cnt += PathFrom(root.Right, sum-root.Val)

	return cnt
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
