package answer

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	curr1, curr2 := headA, headB
//	NodeMap := make(map[*ListNode]struct{})
//	for curr1 != nil {
//		if _, ok := NodeMap[curr1]; ok {
//			return curr1
//		} else {
//			NodeMap[curr1] = struct{}{}
//		}
//		curr1 = curr1.Next
//	}
//	for curr2 != nil {
//		if _, ok := NodeMap[curr2]; ok {
//			return curr2
//		} else {
//			NodeMap[curr1] = struct{}{}
//		}
//		curr2 = curr2.Next
//	}
//	return nil
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
