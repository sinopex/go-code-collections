// 617 合并二叉树
// https://leetcode-cn.com/problems/merge-two-binary-trees/
package leetcode

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	val := 0
	var t1Left, t1Right, t2Left, t2Right *TreeNode
	if t1 != nil {
		val += t1.Val
		t1Left, t1Right = t1.Left, t1.Right
	}
	if t2 != nil {
		val += t2.Val
		t2Left, t2Right = t2.Left, t2.Right
	}

	return &TreeNode{
		Val:   val,
		Left:  mergeTrees(t1Left, t2Left),
		Right: mergeTrees(t1Right, t2Right),
	}
}
