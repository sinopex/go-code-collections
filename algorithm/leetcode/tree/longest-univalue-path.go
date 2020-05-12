// 687 最长同值路径
// https://leetcode-cn.com/problems/longest-univalue-path/
package tree

func longestUnivaluePath(root *TreeNode) int {
	var maxPath int
	dfs(root, 0, &maxPath)
	return maxPath
}

// f 为父节点的值
func dfs(root *TreeNode, f int, maxPath *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, root.Val, maxPath)  // 左树最长路径
	r := dfs(root.Right, root.Val, maxPath) // 右树最长路径
	*maxPath = max(r+l, *maxPath)           // 路径是否变长
	if f == root.Val {                      // 与父节点的值相等时，并且选择l或r中最长路径
		return max(l, r) + 1
	}
	// 与父节点不相等，则这是一个新的起始节点
	// 重新纳入最长路径计算
	return 0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
