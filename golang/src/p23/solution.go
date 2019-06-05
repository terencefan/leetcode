package main

import (
	"container/heap"
)

type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len() - 1]
	*h = (*h)[:h.Len() - 1]
	return r
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := make(Heap, 0)

	for _, node := range lists {
		if node != nil {
			heap.Push(&h, node)
		}
	}

	root := &ListNode{-1, nil}
	iter := root

	for h.Len() > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		iter.Next = node
		iter = iter.Next
	}

	return root.Next
}
