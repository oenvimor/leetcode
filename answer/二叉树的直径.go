package answer

func DiameterOfBinaryTree(root *TreeNode) int {
	longest := 0
	length := getLength(root, &longest)
	if root.Left == nil || root.Right == nil {
		return longest
	}
	return Max(longest, length)
}

func getLength(node *TreeNode, longest *int) int {
	if node == nil {
		return 0
	}
	left := getLength(node.Left, longest)
	right := getLength(node.Right, longest)
	length := left + right
	if length > *longest {
		*longest = length
	}
	return Max(left, right) + 1
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
