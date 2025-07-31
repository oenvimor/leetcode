package answer

import (
	"container/heap"
)

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 定义一个小顶堆
	h := &ListNodeHeap{}
	// 初始化
	heap.Init(h)
	// 将每一个链表的头节点放入堆中
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		// 拿到最顶也是值最小的 node
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}
