package p767

import "container/heap"

type ByteCount struct {
	count int
	char  byte
}

type ByteHeap []ByteCount

// Len implements heap.Interface.
func (b *ByteHeap) Len() int {
	return len(*b)
}

// Less implements heap.Interface.
func (b *ByteHeap) Less(i int, j int) bool {
	return (*b)[i].count > (*b)[j].count
}

// Pop implements heap.Interface.
func (b *ByteHeap) Pop() any {
	r := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return r
}

// Push implements heap.Interface.
func (b *ByteHeap) Push(x any) {
	*b = append(*b, x.(ByteCount))
}

// Swap implements heap.Interface.
func (b *ByteHeap) Swap(i int, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func reorganizeString2(s string) string {
	if len(s) < 2 {
		return s
	}

	var h = make(ByteHeap, 26)
	for i := range 26 {
		h[i] = ByteCount{0, byte(i + 'a')}
	}

	for _, c := range s {
		h[c-'a'].count++
	}

	heap.Init(&h)

	var bytes = make([]byte, len(s))
	for i := range len(s) {
		top := heap.Pop(&h).(ByteCount)
		if i == 0 || top.char != bytes[i-1] {
			bytes[i] = top.char
			top.count--
		} else if h.Len() == 0 {
			return ""
		} else {
			sec := heap.Pop(&h).(ByteCount)
			if sec.count <= 0 {
				return ""
			}
			bytes[i] = sec.char
			sec.count--
			heap.Push(&h, sec)
		}
		heap.Push(&h, top)
	}
	return string(bytes)
}
