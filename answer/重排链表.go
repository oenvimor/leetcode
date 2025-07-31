package answer

func ReorderList(head *ListNode) {
	// 快慢指针法找到链表中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 翻转后半部分链表
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// 合并两个链表
	front, back := head, prev
	for back != nil {
		tmp1 := front.Next
		tmp2 := back.Next

		front.Next = back
		back.Next = tmp1

		front = tmp1
		back = tmp2
	}
}
