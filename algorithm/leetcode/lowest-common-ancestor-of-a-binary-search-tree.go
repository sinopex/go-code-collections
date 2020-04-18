// 235 二叉搜索树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
