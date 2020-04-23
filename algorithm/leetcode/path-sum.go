// 112 路径总和
// https://leetcode-cn.com/problems/path-sum/
package leetcode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	queues := make([]*TreeNode, 0)
	queues = append(queues, root)
	n := len(queues)
	for n > 0 {
		for i := 0; i < n; i++ {
			// 已经到了叶子节点
			if queues[i].Left == nil && queues[i].Right == nil {
				if queues[i].Val == sum {
					return true
				}
			}

			if queues[i].Left != nil {
				queues[i].Left.Val += queues[i].Val
				queues = append(queues, queues[i].Left)
			}

			if queues[i].Right != nil {
				queues[i].Right.Val += queues[i].Val
				queues = append(queues, queues[i].Right)
			}
		}
		queues = queues[n:]
		n = len(queues)
	}

	return false
}

func hasPathSumDfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
