package answer

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var preOrder func(root *TreeNode)
	res := make([]int, 0)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return res
}
