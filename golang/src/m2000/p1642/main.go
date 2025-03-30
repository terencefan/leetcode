package p1642

import "container/heap"

type IntHeap []int

// Len implements heap.Interface.
func (i *IntHeap) Len() int {
	return len(*i)
}

// Less implements heap.Interface.
func (h *IntHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Pop implements heap.Interface.
func (i *IntHeap) Pop() any {
	last := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return last
}

// Push implements heap.Interface.
func (i *IntHeap) Push(x any) {
	*i = append(*i, x.(int))
}

// Swap implements heap.Interface.
func (h *IntHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	if len(heights) == 0 {
		return 0
	}
	h := make(IntHeap, 0)

	for i := 1; i < len(heights); i++ {
		prev, height := heights[i-1], heights[i]
		if height <= prev {
			continue
		}

		if ladders > 0 {
			heap.Push(&h, height-prev)
			ladders--
			continue
		} else {
			diff := height - prev

			minLadder := 1 << 31
			if h.Len() > 0 {
				minLadder = h[0]
			}

			if diff > minLadder {
				heap.Pop(&h)
				heap.Push(&h, diff)
				diff = minLadder
			}

			if bricks >= diff {
				bricks -= diff
			} else {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}
