package main

import "container/heap"

type IntHeap []int

type MedianFinder struct {
	l *IntHeap
	r *IntHeap
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	_ = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	var l = len(*h)
	defer func() {
		*h = (*h)[:l-1]
	}()
	return (*h)[l-1]
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.l, num)
}

func (this *MedianFinder) FindMedian() float64 {
	return 1.5
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
