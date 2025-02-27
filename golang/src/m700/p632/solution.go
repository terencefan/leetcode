package main

import (
	"container/heap"
	"fmt"
)

type HeapNode struct {
	k, index, val int
}

type Heap []*HeapNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*HeapNode))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

type Solver struct {
	h          Heap
	currentMax int
}

func newHeap() *Solver {
	return &Solver{
		h:          make([]*HeapNode, 0),
		currentMax: -(1 << 32),
	}
}

func (s *Solver) Push(n *HeapNode) {
	if n.val > s.currentMax {
		s.currentMax = n.val
	}
	heap.Push(&s.h, n)
}

func (s *Solver) Pop() *HeapNode {
	return heap.Pop(&s.h).(*HeapNode)
}

func smallestRange(nums [][]int) []int {
	h := newHeap()
	for k := range nums {
		if len(nums[k]) == 0 {
			return []int{0, 0} // this won't happen
		}
		h.Push(&HeapNode{k, 0, nums[k][0]})
	}

	delta := 1<<32 - 1

	r := make([]int, 2)
	for {
		node := h.Pop()
		if h.currentMax-node.val < delta {
			r[0], r[1] = node.val, h.currentMax
			delta = h.currentMax - node.val
		}
		if node.index == len(nums[node.k])-1 {
			break
		}
		h.Push(&HeapNode{node.k, node.index + 1, nums[node.k][node.index+1]})
	}
	return r
}

func main() {
	r := smallestRange([][]int{
		{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30},
	})
	fmt.Println(r)
}
