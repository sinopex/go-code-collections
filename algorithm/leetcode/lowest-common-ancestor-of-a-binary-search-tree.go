// 235 二叉搜索树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// p小，q大
	if p.Val > q.Val {
		p, q = q, p
	}
	return _lowestCommonAncestor(root, p, q)
}

func _lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// p,q 在root的两边，那么顶点就是最近公共祖先
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}

	// root在右侧，则往左下沉
	if root.Val > p.Val && root.Val > q.Val {
		return _lowestCommonAncestor(root.Left, p, q)
	}

	// root在左侧，则往右下沉
	if root.Val < p.Val && root.Val < q.Val {
		return _lowestCommonAncestor(root.Right, p, q)
	}

	return nil
}
