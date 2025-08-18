package answer

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	curr := dummy

	incr := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val
		digit := sum + incr
		if digit >= 10 {
			digit = digit % 10
			incr = 1
		} else {
			incr = 0
		}
		curr.Next = &ListNode{Val: digit}
		l1 = l1.Next
		l2 = l2.Next
		curr = curr.Next
	}
	for l2 != nil {
		if l2.Val+incr >= 10 {
			curr.Next = &ListNode{Val: (l2.Val + incr) % 10}
			incr = 1
			l2 = l2.Next
			curr = curr.Next
		} else {
			curr.Next = &ListNode{Val: l2.Val + incr}
			incr = 0
			l2 = l2.Next
			curr = curr.Next
		}
	}
	for l1 != nil {
		if l1.Val+incr >= 10 {
			curr.Next = &ListNode{Val: (l1.Val + incr) % 10}
			incr = 1
			l1 = l1.Next
			curr = curr.Next
		} else {
			curr.Next = &ListNode{Val: l1.Val + incr}
			incr = 0
			l1 = l1.Next
			curr = curr.Next
		}
	}
	if incr == 1 {
		curr.Next = &ListNode{Val: incr}
	}
	return dummy.Next
}
