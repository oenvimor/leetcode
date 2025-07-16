package answer

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	// 创建一个虚拟头节点
	dummy := &ListNode{Next: head}
	// 创建一个前节点
	prev := dummy
	// 计算间隔
	bet := right - left
	// 定义起始节点
	var start, end *ListNode
	// 找到 start 节点的前一个节点
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	start = prev.Next
	end = start
	for i := 0; i < bet; i++ {
		end = end.Next
	}
	// 记录翻转组的下一节点
	next := end.Next
	// 获取翻转之后的新头节点
	newHead := reverse(start, end.Next)
	// 将前节点与新头节点接上
	prev.Next = newHead
	// 将尾节点与下一个节点接上
	start.Next = next
	// 返回头节点
	return dummy.Next
}

func reverse(start, end *ListNode) *ListNode {
	prev := &ListNode{}
	curr := start
	for curr != end {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
