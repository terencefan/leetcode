package p295

import "container/heap"

type IntHeap []int

// Len implements heap.Interface.
func (h *IntHeap) Len() int {
	return len(*h)
}

// Less implements heap.Interface.
func (h *IntHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Pop implements heap.Interface.
func (h *IntHeap) Pop() any {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

// Push implements heap.Interface.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h *IntHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

type MedianFinder struct {
	minHeap, maxHeap IntHeap
	count            int
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: make(IntHeap, 0),
		maxHeap: make(IntHeap, 0),
	}
}

func (finder *MedianFinder) AddNum(num int) {
	if len(finder.minHeap) > 0 && num >= finder.minHeap[0] {
		heap.Push(&finder.minHeap, num)
		for len(finder.minHeap)-len(finder.maxHeap) > 1 {
			num := heap.Pop(&finder.minHeap).(int)
			heap.Push(&finder.maxHeap, -num)
		}
	} else {
		heap.Push(&finder.maxHeap, -num)
		for len(finder.minHeap)-len(finder.maxHeap) < 0 {
			num := -heap.Pop(&finder.maxHeap).(int)
			heap.Push(&finder.minHeap, num)
		}
	}
	finder.count++
}

func (finder *MedianFinder) FindMedian() float64 {
	if finder.count == 0 {
		return 0
	} else if finder.count%2 == 0 {
		return float64(finder.minHeap[0]-finder.maxHeap[0]) / 2
	} else {
		return float64(finder.minHeap[0])
	}
}
