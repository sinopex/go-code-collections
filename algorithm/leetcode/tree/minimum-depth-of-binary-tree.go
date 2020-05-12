// 111 二叉树的最小深度
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right != nil {
		l := minDepth(root.Left)
		r := minDepth(root.Right)

		if l > r {
			return r + 1
		}

		return l + 1
	}

	if root.Left != nil {
		return minDepth(root.Left) + 1
	}

	return minDepth(root.Right) + 1
}

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queues := []*TreeNode{root}
	n, level := len(queues), 1
	for n > 0 {
		for i := 0; i < n; i++ {
			if queues[i].Left == nil && queues[i].Right == nil {
				return level
			}
			if queues[i].Left != nil {
				queues = append(queues, queues[i].Left)
			}
			if queues[i].Right != nil {
				queues = append(queues, queues[i].Right)
			}
		}
		level++
		queues = queues[n:]
		n = len(queues)
	}

	return level
}
