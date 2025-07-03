package pkg

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 创建一个虚拟头节点
	dummy := &ListNode{Next: head}
	// 创建一个临时节点，用于移动
	pre := dummy
	// 进入循环检查链表是否还有 k 个节点，若没有 k 个就直接返回头节点
	for {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		// 记录下一组节点的开头（因为当前组节点即将翻转，之后就不再好操作获取下一组节点的开头了）
		next := tail.Next
		// 将循环获取到的一组节点进行翻转，并返回翻转后的头节点
		start := pre.Next
		newHead := Reverse(start, next)
		// 将临时节点指向新的头节点
		pre.Next = newHead
		// 将上一组节点的尾节点指向下一组节点的开始节点
		start.Next = next
		// 将临时节点移动到下一组开始节点的上一个节点
		pre = start
	}
}

// Reverse 翻转链表（左闭右开）
func Reverse(head, end *ListNode) *ListNode {
	// 创建一个临时节点，用于移动
	var prev *ListNode = nil
	// 创建一个当前节点
	curr := head
	// 如果当前节点没有到下一组的头节点就一直循环
	for curr != end {
		// 记录下一个节点
		next := curr.Next
		// 改变当前节点的指向
		curr.Next = prev
		// 将 prev 节点先后移动
		prev = curr
		// 改变当前节点为下一个节点
		curr = next
	}
	// 返回头节点，此时头节点是翻转之前的最后一个节点
	return prev
}
