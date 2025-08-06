package answer

func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return dummy.Next
}
