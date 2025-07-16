package answer

//func HasCycle(head *ListNode) bool {
//	// 创建一个当前接节点,用于移动
//	curr := &ListNode{Next: head}
//	// 创建一个 map 用于记录节点是否出现过
//	valMap := make(map[*ListNode]struct{})
//	// 进入循环遍历
//	for curr != nil {
//		// 如果当前节点出现过则返回 true
//		if _, found := valMap[curr]; found {
//			return true
//		}
//		// 没有出现过就加入 map
//		valMap[curr] = struct{}{}
//		curr = curr.Next
//	}
//	return false
//}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
