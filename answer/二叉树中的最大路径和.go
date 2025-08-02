package answer

func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 31
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := Max(0, dfs(node.Left))
		right := Max(0, dfs(node.Right))
		current := left + right + node.Val
		maxSum = Max(maxSum, current)
		return node.Val + Max(left, right)
	}
	dfs(root)
	return maxSum
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
