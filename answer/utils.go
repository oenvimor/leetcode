package answer

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrToList(arr []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, v := range arr {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head.Next
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.Next
	}
}
