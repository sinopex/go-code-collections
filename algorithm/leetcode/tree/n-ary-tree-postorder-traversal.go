// 590 N叉树的后序遍历
// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
package tree

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	for _, node := range root.Children {
		if ret := postorder(node); len(ret) > 0 {
			res = append(res, ret...)
		}
	}
	res = append(res, root.Val)
	return res
}

func postorderBFS(root *Node) []int {
	if root == nil {
		return []int{}
	}

	slice := []*Node{root}
	result := []int{}
	for len(slice) != 0 {
		slice, root = slice[0:len(slice)-1], slice[len(slice)-1]
		result = append([]int{root.Val}, result...)
		for _, value := range root.Children {
			slice = append(slice, value)
		}
	}
	return result
}
