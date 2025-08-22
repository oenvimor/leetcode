package answer

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	if left == -1 {
		return -1
	}
	right := height(node.Right)
	if right == -1 {
		return -1
	}
	if Abs(left, right) > 1 {
		return -1
	}
	return Max(left, right) + 1
}

func Abs(a, b int) int {
	if (a - b) < 0 {
		return b - a
	}
	return a - b
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
