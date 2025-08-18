package answer

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	idx := 0
	for i, v := range inorder {
		if v == rootVal {
			idx = i
			break
		}
	}
	root.Left = BuildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = BuildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
