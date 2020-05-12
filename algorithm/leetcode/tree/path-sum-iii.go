// 437 路径总和
// https://leetcode-cn.com/problems/path-sum-iii/
package tree

func pathSum(root *TreeNode, sum int) int {
	var c int
	var v = make(map[int]int)
	v[0] = 1
	_pathSum(root, sum, 0, &c, v)
	return c
}

func _pathSum(root *TreeNode, sum, ps int, c *int, v map[int]int) {
	if root == nil {
		return
	}

	ps += root.Val
	*c += v[ps-sum]
	v[ps] += 1
	_pathSum(root.Left, sum, ps, c, v)
	_pathSum(root.Right, sum, ps, c, v)
	v[ps] -= 1
}

// func pathSum(root *TreeNode, sum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return PathFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
// }
//
// func PathFrom(root *TreeNode, sum int) int {
// 	if root == nil {
// 		return 0
// 	}
//
// 	cnt := 0
//
// 	if root.Val == sum {
// 		cnt++
// 	}
//
// 	cnt += PathFrom(root.Left, sum-root.Val)
// 	cnt += PathFrom(root.Right, sum-root.Val)
//
// 	return cnt
// }
