package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

func print(tree *TreeNode) []int {
	result := make([]int, 0)
	_print(tree, &result)
	return result
}

func _print(tree *TreeNode, result *[]int) {
	if tree.Left != nil {
		_print(tree.Left, result)
	}
	*result = append(*result, tree.Val)
	if tree.Right != nil {
		_print(tree.Right, result)
	}
}
