package answer

func WidthOfBinaryTree(root *TreeNode) int {
	type posNode struct {
		node *TreeNode
		pos  int
	}
	maxWidth := 0
	queue := []posNode{{root, 0}}
	for len(queue) > 0 {
		size := len(queue)
		left := queue[0].pos
		right := queue[size-1].pos
		width := right - left + 1
		if width > maxWidth {
			maxWidth = width
		}
		newQueue := make([]posNode, 0)
		for _, np := range queue {
			if np.node.Left != nil {
				newQueue = append(newQueue, posNode{np.node.Left, 2 * (np.pos - left)})
			}
			if np.node.Right != nil {
				newQueue = append(newQueue, posNode{np.node.Right, 2*(np.pos-left) + 1})
			}
		}
		queue = newQueue
	}
	return maxWidth
}
