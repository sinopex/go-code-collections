// 572 另一个树的子树
// https://leetcode-cn.com/problems/subtree-of-another-tree/
package tree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if _isSampleTree(s, t) {
		return true
	}

	if s == nil {
		return false
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// 比较两个树是否相等
func _isSampleTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if t == nil || s == nil {
		return false
	}

	if t.Val != s.Val {
		return false
	}

	return _isSampleTree(s.Left, t.Left) && _isSampleTree(s.Right, t.Right)
}

// 摘抄网上同学的精简例子
// 链接：https://leetcode-cn.com/problems/subtree-of-another-tree/solution/dfs-by-linbingyuan-13/// func isSubtree(s *TreeNode, t *TreeNode) bool {
// 	if s == nil {
// 		return false
// 	}
// 	return isSubtree(s.Left, t) || isSameTree(s, t) || isSubtree(s.Right, t)
// }
//
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil || q == nil {
// 		return p == q
// 	}
// 	return isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right)
// }
//
