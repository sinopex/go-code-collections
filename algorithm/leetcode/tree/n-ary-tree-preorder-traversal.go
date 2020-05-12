// 589 N叉树的前序遍历
// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
package tree

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	res := []int{root.Val}

	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res
}

func preorderBFS(root *Node) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	queues := []*Node{root}
	for len(queues) > 0 {
		node := queues[0]
		queues = queues[1:]
		res = append(res, node.Val)

		var tmpNode []*Node
		for _, node := range node.Children {
			tmpNode = append(tmpNode, node)
		}
		queues = append(tmpNode, queues...)
	}

	return res
}
