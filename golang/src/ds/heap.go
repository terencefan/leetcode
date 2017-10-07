package ds

import (
	"container/heap"
)

const HEAPMAX int8 = 0
const HEAPMIN int8 = 1

type IntArray struct {
	arr []int
	tab int8
}

func (h IntArray) Len() int { return len(h.arr) }

func (h IntArray) Less(i, j int) bool {
	if h.tab == HEAPMIN {
		return h.arr[i] < h.arr[j]
	} else {
		return h.arr[i] > h.arr[j]
	}
}

func (h IntArray) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *IntArray) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

func (h *IntArray) Pop() interface{} {
	var l = len(h.arr)
	defer func() {
		h.arr = h.arr[:l-1]
	}()
	return h.arr[l-1]
}

func (h *IntArray) Init(tab int8, x ...int) {
	h.tab = tab
	h.arr = append(h.arr, x...)
}

type IntHeap struct {
	arr *IntArray
}

func NewIntHeap(tab int8, x ...int) (h *IntHeap) {
	var arr = &IntArray{}
	arr.Init(tab, x...)
	heap.Init(arr)
	h = &IntHeap{arr: arr}
	return
}

func (this *IntHeap) Push(x int) {
	heap.Push(this.arr, x)
}

func (this *IntHeap) Pop() (r int) {
	r, _ = heap.Pop(this.arr).(int)
	return
}
