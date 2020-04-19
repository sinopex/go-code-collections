// 107 二叉树的层次遍历II
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queues := make([]*TreeNode, 0)
	ret := make([][]int, 0)

	queues = append(queues, root)
	n := len(queues)

	for n > 0 {
		t := make([]int, 0, n)
		for i := 0; i < n; i++ {
			t = append(t, queues[i].Val)
			if queues[i].Left != nil {
				queues = append(queues, queues[i].Left)
			}
			if queues[i].Right != nil {
				queues = append(queues, queues[i].Right)
			}
		}

		if len(t) > 0 {
			ret = append([][]int{t}, ret...)
		}

		queues = queues[n:]
		n = len(queues)
	}

	return ret
}

func levelOrderBottomDFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	_levelOrderBottomDFS(root, &ret, 0)

	resultLength := len(ret)
	left, right := 0, resultLength-1
	for left < right {
		ret[left], ret[right] = ret[right], ret[left]
		left++
		right--
	}

	return ret
}

func _levelOrderBottomDFS(root *TreeNode, ret *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*ret) > level {
		(*ret)[level] = append((*ret)[level], root.Val)
	} else {
		*ret = append(*ret, []int{root.Val})
	}

	_levelOrderBottomDFS(root.Left, ret, level+1)
	_levelOrderBottomDFS(root.Right, ret, level+1)
}
