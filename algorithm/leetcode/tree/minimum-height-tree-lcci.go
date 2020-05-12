// 最小高度数
// https://leetcode-cn.com/problems/minimum-height-tree-lcci/
package tree

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[n/2],
		Left:  sortedArrayToBST(nums[0 : n/2]),
		Right: sortedArrayToBST(nums[n/2+1:]),
	}
}
