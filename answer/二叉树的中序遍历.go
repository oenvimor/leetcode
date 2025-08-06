package answer

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var inOrder func(*TreeNode)
	res := make([]int, 0)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return res
}
