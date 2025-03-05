package heap

func children(i int) (int, int) {
	return i<<1 + 1, i<<1 + 2
}

func parent(i int) int {
	return (i - 1) >> 1
}

type Heap struct {
	data     []any
	comparer func(i, j int) bool
}

func (h *Heap) len() int {
	return len(h.data)
}

func (h *Heap) heapify(i int) {
	for i > 0 {
		p := parent(i)
		if h.comparer(i, p) {
			h.swap(i, p)
			i = p
		} else {
			break
		}
	}

	k := i
	for {
		l, r := children(i)
		if l < h.len() && h.comparer(l, i) {
			i = l
		}
		if r < h.len() && h.comparer(r, i) {
			i = r
		}
		if k == i {
			break
		}
		h.swap(i, k)
		i = k
	}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Peek() any {
	if len(h.data) == 0 {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Pop() any {
	if len(h.data) == 0 {
		return nil
	}
	r := h.Peek()
	h.swap(0, h.len()-1)
	h.data = h.data[:h.len()-1]
	h.heapify(0)
	return r
}

func (h *Heap) Push(v any) {
	h.data = append(h.data, v)
	h.heapify(h.len() - 1)
}

func Slice(data []any, comparer func(i, j int) bool) *Heap {
	h := &Heap{data, comparer}
	for i := range data {
		h.heapify(i)
	}
	return h
}
