package main

import (
	"fmt"
)

type MinHeap []int

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) rise(i int) int {
	var p = parent(i)

	if (*h)[i] < (*h)[p] {
		h.swap(i, p)
		return p
	}
	return 0
}

func (h *MinHeap) sink(i int) int {
	if i >= len(*h)/2 {
		return -1
	}
	var l, r = children(i)

	if r < len(*h) {
		if (*h)[l] < (*h)[r] {
			if (*h)[l] < (*h)[i] {
				h.swap(i, l)
				return l
			}
		} else {
			if (*h)[r] < (*h)[i] {
				h.swap(i, r)
				return r
			}
		}
	} else {
		if (*h)[l] < (*h)[i] {
			h.swap(i, l)
			return l
		}
	}
	return -1
}

func (h *MinHeap) Push(x int) {
	*h = append(*h, x)

	var i = len(*h) - 1

	for i > 0 {
		i = h.rise(i)
	}
}

func (h *MinHeap) Pop() int {
	r := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	var i = 0
	for i >= 0 {
		i = h.sink(i)
	}

	return r
}

type SmallestInfiniteSet struct {
	heap *MinHeap
	m    map[int]int
	min  int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		heap: &MinHeap{},
		m:    make(map[int]int),
		min:  0,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(*this.heap) == 0 || (*this.heap)[0] > this.min {
		this.min++
		return this.min
	}

	var r = this.heap.Pop()
	delete(this.m, r)
	return r
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.m[num]; !ok && num <= this.min {
		this.m[num] = 0
		this.heap.Push(num)
	}
}

func main() {
	var h = MinHeap{}

	h.Push(12)
	h.Push(8)
	h.Push(10)
	h.Push(6)
	h.Push(7)
	h.Push(2)

	for len(h) > 0 {
		fmt.Println(h.Pop())
	}
}
