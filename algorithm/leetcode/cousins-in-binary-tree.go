// 993 二叉树的堂兄弟节点
// https://leetcode-cn.com/problems/cousins-in-binary-tree/
package leetcode

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	queues := make([]*TreeNode, 0)
	queues = append(queues, root)
	n := len(queues)

	for n > 0 {
		var vx, vy bool // x和y在本层是否找到
		var px, py int  // x和有对应的父节点的值
		for i := 0; i < n; i++ {
			if queues[i].Left != nil {
				queues = append(queues, queues[i].Left)
				// 左子树遍历时，有没有找到目标数字
				if queues[i].Left.Val == x {
					vx = true
					px = queues[i].Val
				}
				if queues[i].Left.Val == y {
					vy = true
					py = queues[i].Val
				}

			}
			if queues[i].Right != nil {
				// 右子树遍历时，有没有找到目标数字
				queues = append(queues, queues[i].Right)
				if queues[i].Right.Val == x {
					vx = true
					px = queues[i].Val
				}
				if queues[i].Right.Val == y {
					vy = true
					py = queues[i].Val
				}
			}
		}

		// 如果两个节点不在一层，肯定不是堂兄，到此可以结束
		if (vx && !vy) || (vy && !vx) {
			return false
		}

		//  这一层两个节点全部出现，比较父节点是否为同一个即可
		if vx && vy {
			return px != py
		}

		// 循环下一层
		queues = queues[n:]
		n = len(queues)
	}

	return false
}
