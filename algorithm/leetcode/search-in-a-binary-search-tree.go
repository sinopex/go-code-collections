// 700 二叉搜索树种的搜索
// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	// 充分运用二叉搜索树的特性
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}
