package answer

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	if root.Left != nil || root.Right != nil {
		left := root.Left
		root.Left = root.Right
		root.Right = left
	}
	return root
}
