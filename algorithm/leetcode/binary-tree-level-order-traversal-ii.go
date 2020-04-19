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
