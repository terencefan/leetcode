package p23

import (
	"container/heap"

	"terencefan.com/leetcode/src/utils"
)

type ListNode = utils.ListNode

type ListNodeHeap []*ListNode

// Len implements heap.Interface.
func (l *ListNodeHeap) Len() int {
	return len(*l)
}

// Less implements heap.Interface.
func (l *ListNodeHeap) Less(i int, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}

// Pop implements heap.Interface.
func (l *ListNodeHeap) Pop() any {
	r := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return r
}

// Push implements heap.Interface.
func (l *ListNodeHeap) Push(x any) {
	*l = append(*l, x.(*ListNode))
}

// Swap implements heap.Interface.
func (l *ListNodeHeap) Swap(i int, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	var h = ListNodeHeap{}

	heap.Init(&h)

	for _, head := range lists {
		if head != nil {
			heap.Push(&h, head)
		}
	}

	var dummy = &ListNode{}
	var current = dummy
	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		current.Next = node
		current = current.Next
	}
	return dummy.Next
}
