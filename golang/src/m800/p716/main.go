package p716

import "container/heap"

type IndexVal struct {
	idx, val int
}

type MaxHeap []IndexVal

// Len implements heap.Interface.
func (m *MaxHeap) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MaxHeap) Less(i int, j int) bool {
	if (*m)[i].val == (*m)[j].val {
		return (*m)[i].idx > (*m)[j].idx
	}
	return (*m)[i].val > (*m)[j].val
}

// Pop implements heap.Interface.
func (m *MaxHeap) Pop() any {
	r := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return r
}

// Push implements heap.Interface.
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(IndexVal))
}

// Swap implements heap.Interface.
func (m *MaxHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

type MaxStack struct {
	removed map[int]bool
	h       MaxHeap
	s       []IndexVal
	cnt     int
}

func Constructor() MaxStack {
	return MaxStack{
		removed: map[int]bool{},
		h:       MaxHeap{},
		s:       []IndexVal{},
		cnt:     0,
	}
}

func (this *MaxStack) Push(x int) {
	this.cnt++
	item := IndexVal{this.cnt, x}
	heap.Push(&this.h, item)
	this.s = append(this.s, item)
}

func (this *MaxStack) Top() int {
	for len(this.s) > 0 {
		last := this.s[len(this.s)-1]
		if this.removed[last.idx] {
			delete(this.removed, last.idx)
			this.s = this.s[:len(this.s)-1]
		} else {
			return last.val
		}
	}
	return -1
}

func (this *MaxStack) PeekMax() int {
	for this.h.Len() > 0 && this.removed[this.h[0].idx] {
		item := heap.Pop(&this.h).(IndexVal)
		delete(this.removed, item.idx)
	}
	return this.h[0].val
}

func (this *MaxStack) Pop() int {
	for len(this.s) > 0 {
		last := this.s[len(this.s)-1]
		this.s = this.s[:len(this.s)-1]
		if this.removed[last.idx] {
			delete(this.removed, last.idx)
		} else {
			this.removed[last.idx] = true
			return last.val
		}
	}
	return -1
}

func (this *MaxStack) PopMax() int {
	for this.h.Len() > 0 && this.removed[this.h[0].idx] {
		item := heap.Pop(&this.h).(IndexVal)
		delete(this.removed, item.idx)
	}
	item := heap.Pop(&this.h).(IndexVal)
	this.removed[item.idx] = true
	return item.val
}
