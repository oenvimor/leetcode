package answer

func HasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode, sum int) bool
	dfs = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			return true
		}
		if dfs(node.Left, sum) {
			return true
		}
		if dfs(node.Right, sum) {
			return true
		}
		return false
	}
	return dfs(root, 0)
}
