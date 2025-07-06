package answer

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	// 将根节点放入队列中
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)          // 用于记录当前层节点的数量
		arr := make([]int, 0, size) // 用于记录当前层节点的值
		for i := 0; i < size; i++ {
			node := queue[0]            // 取出第一个节点
			queue = queue[1:]           // 将第一个节点出队
			arr = append(arr, node.Val) // 记录当前节点的值
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, arr) // 记录当前层的节点的值
	}
	return res
}
