package answer

func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 维护一个从左往右遍历的标签
	var leftToRight bool
	// 维护一个队列
	var queue []*TreeNode
	// 维护一个返回的数组
	var ans [][]int
	// 将根节点加入队列
	queue = append(queue, root)
	// 记录层数
	var count int
	// 进入循环
	for len(queue) > 0 {
		// 偶数层从左往右,奇数层从右往左
		if count&1 == 1 {
			leftToRight = false
		} else {
			leftToRight = true
		}
		// 记录当前层节点的数量
		size := len(queue)
		// 创建一个临时切片
		tmp := make([]int, size)
		// 遍历当前层每一个节点
		for i := 0; i < size; i++ {
			// 取出一个节点
			node := queue[0]
			// 将第一个节点出队
			queue = queue[1:]
			// 将节点的值加入数组
			if leftToRight {
				tmp[i] = node.Val
			} else {
				tmp[size-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, tmp)
		count++
	}
	return ans
}
