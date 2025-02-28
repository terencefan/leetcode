package p215

type MinHeap struct {
	heap     []int
	capacity int
	length   int
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) rise(index int) {
	for index > 0 {
		p := parent(index)
		if h.heap[p] > h.heap[index] {
			h.heap[p], h.heap[index] = h.heap[index], h.heap[p]
			index = p
		} else {
			break
		}
	}
}

func (h *MinHeap) sink(index int) {
	for {
		l, r := children(index)
		smaller := index

		if l < h.length && h.heap[l] < h.heap[smaller] {
			smaller = l
		}
		if r < h.length && h.heap[r] < h.heap[smaller] {
			smaller = r
		}
		if smaller == index {
			break
		} else {
			h.heap[index], h.heap[smaller] = h.heap[smaller], h.heap[index]
			index = smaller
		}
	}
}

func (h *MinHeap) Push(num int) {
	if h.length < h.capacity {
		h.heap[h.length] = num
		h.rise(h.length)
		h.length++
	} else if num > h.Peek() {
		h.heap[0] = num
		h.sink(0)
	}
}

func (h *MinHeap) Peek() int {
	return h.heap[0]
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		heap:     make([]int, capacity+1),
		length:   0,
		capacity: capacity,
	}
}

func findKthLargestHeap(nums []int, k int) int {
	var heap = NewMinHeap(k)

	for _, num := range nums {
		heap.Push(num)
	}
	return heap.Peek()
}
