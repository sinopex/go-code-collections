// 872 叶子相似的树
// https://leetcode-cn.com/problems/leaf-similar-trees/
package tree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ret1 := make([]int, 0)
	ret2 := make([]int, 0)
	_leafSimilar(root1, &ret1)
	_leafSimilar(root2, &ret2)

	if len(ret1) != len(ret2) {
		return false
	}

	for k, v := range ret1 {
		if ret2[k] != v {
			return false
		}
	}

	return true
}
func _leafSimilar(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, root.Val)
		return
	}

	_leafSimilar(root.Left, ret)
	_leafSimilar(root.Right, ret)
}
