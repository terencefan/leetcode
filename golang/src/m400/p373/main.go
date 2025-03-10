package p373

import "container/heap"

type IndexVal struct {
	x, y, val int
}

type MinHeap []IndexVal

// Len implements heap.Interface.
func (m *MinHeap) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MinHeap) Less(i int, j int) bool {
	return (*m)[i].val < (*m)[j].val
}

// Pop implements heap.Interface.
func (m *MinHeap) Pop() any {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

// Push implements heap.Interface.
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(IndexVal))
}

// Swap implements heap.Interface.
func (m *MinHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	h := &MinHeap{}

	for j := range min(k, len(nums2)) {
		heap.Push(h, IndexVal{0, j, nums1[0] + nums2[j]})
	}

	var r = make([][]int, 0)
	for k > 0 && len(*h) > 0 {
		d := heap.Pop(h).(IndexVal)
		r = append(r, []int{nums1[d.x], nums2[d.y]})

		if d.x+1 < len(nums1) {
			heap.Push(h, IndexVal{d.x + 1, d.y, nums1[d.x+1] + nums2[d.y]})
		}
		k--
	}
	return r
}
