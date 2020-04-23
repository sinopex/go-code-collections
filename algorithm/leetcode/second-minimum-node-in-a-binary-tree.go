// 671 二叉树中第二小的节点
// https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/
package leetcode

func findSecondMinimumValue(root *TreeNode) int {
	var data = make([]int, 0)

	_findSecondMinimumValue(root, &data)

	if len(data) == 2 {
		return data[1]
	}

	return -1
}
func _findSecondMinimumValue(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	n := len(*data)
	if n == 0 {
		*data = append(*data, root.Val)
	} else if n == 1 {
		if root.Val != (*data)[0] {
			*data = append(*data, root.Val)
		}
	} else {
		// 第二小的需要判断
		if root.Val != (*data)[0] && (*data)[1] > root.Val {
			(*data)[1] = root.Val
		}
	}
	_findSecondMinimumValue(root.Left, data)
	_findSecondMinimumValue(root.Right, data)
}
