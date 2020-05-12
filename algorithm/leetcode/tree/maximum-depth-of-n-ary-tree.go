// 559 N叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
package tree

func maxDepthN(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	queues := []*Node{root}
	n := len(queues)
	for n > 0 {
		depth++
		for i := 0; i < n; i++ {
			for _, node := range queues[i].Children {
				queues = append(queues, node)
			}
		}
		queues = queues[n:]
		n = len(queues)
	}

	return depth
}
