package answer

func PathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int

	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		// 选择当前节点
		path = append(path, node.Val)
		sum += node.Val

		// 如果是叶子节点并且满足条件
		if node.Left == nil && node.Right == nil && sum == targetSum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		// 继续递归左右子树
		dfs(node.Left, sum)
		dfs(node.Right, sum)

		// 回溯：撤销选择
		path = path[:len(path)-1]
	}

	dfs(root, 0)
	return res
}
