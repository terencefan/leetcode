package p2530

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	if len(nums) == 0 {
		return 0
	}

	h := IntHeap{}
	for _, num := range nums {
		heap.Push(&h, num)
	}

	r := int64(0)
	for range k {
		x := heap.Pop(&h).(int)
		r += int64(x)

		if x%3 == 0 {
			heap.Push(&h, x/3)
		} else {
			heap.Push(&h, x/3+1)
		}
	}
	return r
}
