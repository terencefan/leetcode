package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

type IntMinHeap struct {
	IntHeap
}

type IntMaxHeap struct {
	IntHeap
}

func (h IntHeap) Len() int      { return len(h) }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() (r interface{}) {
	var l = len(*h)
	r = (*h)[l-1]
	*h = (*h)[:l-1]
	return
}

func (h IntMaxHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
func (h IntMinHeap) Less(i, j int) bool { return h.IntHeap[i] < h.IntHeap[j] }

func (h IntHeap) peek() int {
	return h[0]
}

type MedianFinder struct {
	min IntMinHeap
	max IntMaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		min: IntMinHeap{make(IntHeap, 0)},
		max: IntMaxHeap{make(IntHeap, 0)},
	}
}

func (f *MedianFinder) AddNum(num int) {
	if f.max.Len() == 0 {
		heap.Push(&f.max, num)
		return
	}

	pivot := f.max.peek()
	if num < pivot {
		heap.Push(&f.max, num)
		if f.max.Len()-f.min.Len() > 1 {
			heap.Push(&f.min, heap.Pop(&f.max))
		}
	} else {
		heap.Push(&f.min, num)
		if f.max.Len() < f.min.Len() {
			heap.Push(&f.max, heap.Pop(&f.min))
		}
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.max.Len() == 0 {
		return 0.0
	}
	if f.max.Len() == f.min.Len() {
		return float64((f.max.peek())+f.min.peek()) / 2
	} else {
		return float64(f.max.peek())
	}
}

func main() {
	f := Constructor()
	fmt.Println(f.FindMedian())
	f.AddNum(1)
	fmt.Println(f.FindMedian())
	f.AddNum(2)
	fmt.Println(f.FindMedian())
	f.AddNum(1)
	fmt.Println(f.FindMedian())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
