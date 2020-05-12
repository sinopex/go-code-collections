// 637 二叉树的层平均值
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
package tree

func averageOfLevels(root *TreeNode) []float64 {
	ret := make([]float64, 0)
	sum := make([]int, 0)
	_averageOfLevels(root, &ret, &sum, 0)
	for i, v := range ret {
		(ret)[i] = v / float64(sum[i])
	}

	return ret
}

func _averageOfLevels(root *TreeNode, ret *[]float64, sum *[]int, level int) {
	if root == nil {
		return
	}

	// 当前层的累计和已经存在，则累加
	if len(*ret) > level {
		(*ret)[level] += float64(root.Val)
		(*sum)[level] += 1
	} else {
		// 当前层的累计不存在，则新增
		*ret = append(*ret, float64(root.Val))
		*sum = append(*sum, 1)
	}

	_averageOfLevels(root.Left, ret, sum, level+1)
	_averageOfLevels(root.Right, ret, sum, level+1)
}
