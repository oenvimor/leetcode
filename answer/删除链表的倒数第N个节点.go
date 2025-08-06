package answer

//func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
//	curr1 := head
//	num := 0
//	for curr1 != nil {
//		curr1 = curr1.Next
//		num++
//	}
//	idx := num - n
//	dummy := &ListNode{Next: head}
//	curr2 := dummy
//	for i := 0; i < idx; i++ {
//		curr2 = curr2.Next
//	}
//	curr2.Next = curr2.Next.Next
//	return dummy.Next
//}

// RemoveNthFromEnd 快慢指针法
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
