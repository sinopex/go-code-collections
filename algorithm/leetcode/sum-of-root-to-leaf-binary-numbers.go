// 1022 从根到叶的二进制之和
// https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/
package leetcode

func sumRootToLeaf(root *TreeNode) int {
	total := 0
	_sumRootToLeaf(root, &total, 0)
	return total
}
func _sumRootToLeaf(root *TreeNode, total *int, start int) {
	if root == nil {
		return
	}

	if root.Val == 0 {
		start = start << 1
	} else {
		start = start<<1 + 1
	}

	if root.Left == nil && root.Right == nil {
		*total += start % (1e9 + 7)
	}

	_sumRootToLeaf(root.Left, total, start)
	_sumRootToLeaf(root.Right, total, start)
}
