// 669 修剪二叉搜索树
// https://leetcode-cn.com/problems/trim-a-binary-search-tree/
package leetcode

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	for {
		if root == nil {
			break
		}

		if root.Val >= L && root.Val <= R {
			break
		}

		// 顶点在左侧
		for root != nil && root.Val < L {
			root = root.Right
		}

		// 顶点在右侧
		for root != nil && root.Val > R {
			root = root.Left
		}
	}
	// 没有找到合适的节点
	if root == nil {
		return nil
	}

	// 剪左边大的
	p := root
	for p.Left != nil {
		if p.Left.Val >= L {
			p = p.Left
		} else {
			p.Left = p.Left.Right
		}
	}
	// 剪右边大的
	p = root
	for p.Right != nil {
		if p.Right.Val <= R {
			p = p.Right
		} else {
			p.Right = p.Right.Left
		}
	}

	return root
}
