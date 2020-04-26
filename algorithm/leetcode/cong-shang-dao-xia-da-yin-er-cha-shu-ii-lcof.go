// 面试题32-II 从上到下打印二叉树II
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
package leetcode

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)
	_levelOrderBottomDFS(root, &res, 0)
	return res
}

func _levelOrder(root *TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*res) > level {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		*res = append(*res, []int{root.Val})
	}
	level++
	_levelOrder(root.Left, res, level)
	_levelOrder(root.Right, res, level)
	level--
}
