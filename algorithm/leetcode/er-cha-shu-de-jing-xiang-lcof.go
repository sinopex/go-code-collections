// 二叉树的镜像
// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
package leetcode

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l, r := root.Right, root.Left

	root.Left = mirrorTree(l)
	root.Right = mirrorTree(r)
	return root
}
