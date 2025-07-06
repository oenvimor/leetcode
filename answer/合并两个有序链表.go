package answer

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个虚拟头节点
	dummy := &ListNode{}
	// 创建一个当前节点的指针
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return dummy.Next
}
