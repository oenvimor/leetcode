package answer

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode // 初始化头节点
	curr := head
	for curr != nil {
		next := curr.Next // 存储下一个节点
		curr.Next = prev  // 改变当前节点的指向
		prev = curr       // 使上一节点改变为当前节点
		curr = next       // 改变当前节点为下一个节点
	}
	return prev
}
