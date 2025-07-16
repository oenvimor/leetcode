package answer

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ArrToList 数组转链表
func ArrToList(arr []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, v := range arr {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head.Next
}

// PrintList 打印链表
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.Next
	}
}

// ArrToCycleList 数组转循环链表
func ArrToCycleList(nums []int, pos int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	p := dummy
	for i, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
		if i == len(nums)-1 && pos != -1 {
			for j := 0; j < pos+1; j++ {
				p = p.Next
			}
			curr.Next = p
		}
	}
	return dummy.Next
}
