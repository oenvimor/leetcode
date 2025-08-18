package answer

func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := &ListNode{Next: head}, &ListNode{Next: head}
	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
