// 653 两数之和IV - 输入 BST
// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
package tree

func findTarget(root *TreeNode, k int) bool {
	var history = make(map[int]struct{})
	return _findTarget(root, k, &history)
}
func _findTarget(root *TreeNode, k int, history *map[int]struct{}) bool {
	if root == nil {
		return false
	}

	// 已经找到，终止继续查找
	if _, ok := (*history)[k-root.Val]; ok {
		return true
	}
	(*history)[root.Val] = struct{}{}

	return _findTarget(root.Left, k, history) || _findTarget(root.Right, k, history)
}
