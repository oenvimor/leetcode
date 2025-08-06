package answer

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	heap := []*TreeNode{root}
	res := make([]int, 0)
	for len(heap) != 0 {
		size := len(heap)
		tmp := heap
		res = append(res, tmp[len(tmp)-1].Val)
		for _, node := range tmp {
			if node.Left != nil {
				heap = append(heap, node.Left)
			}
			if node.Right != nil {
				heap = append(heap, node.Right)
			}
		}
		heap = heap[size:]
	}
	return res
}
