package answer

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			for head != slow {
				head, slow = head.Next, slow.Next
			}
			return head
		}
	}
	return nil
}
