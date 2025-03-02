package p2462

type Worker struct {
	indx int
	cost int
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

type MinHeap struct {
	workers []Worker
	count   int
}

func (h *MinHeap) swap(i, j int) {
	h.workers[i], h.workers[j] = h.workers[j], h.workers[i]
}

func (h *MinHeap) less(i, j int) bool {
	a, b := h.workers[i], h.workers[j]
	if a.cost == b.cost {
		return a.indx < b.indx
	}
	return a.cost < b.cost
}

func (h *MinHeap) Push(x Worker) {
	h.workers[h.count] = x
	h.count++

	i := h.count - 1
	for i > 0 {
		p := parent(i)
		if h.less(i, p) {
			h.swap(i, p)
			i = p
		} else {
			break
		}
	}
}

func (h *MinHeap) Pop() Worker {
	worker := h.workers[0]

	h.count--
	h.swap(0, h.count)

	i := 0
	for i < h.count {
		l, r := children(i)

		s := i
		if l < h.count && h.less(l, i) {
			s = l
		}
		if r < h.count && h.less(r, s) {
			s = r
		}
		if s == i {
			break
		} else {
			h.swap(i, s)
			i = s
		}
	}

	return worker
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)

	var heap = MinHeap{
		workers: make([]Worker, n),
		count:   0,
	}

	l, r := candidates, n-candidates-1
	if l > r {
		for i, cost := range costs {
			heap.Push(Worker{i, cost})
		}
	} else {
		for i := range candidates {
			heap.Push(Worker{i, costs[i]})
			heap.Push(Worker{n - i - 1, costs[n-i-1]})
		}
	}

	var res = int64(0)
	for range k {
		worker := heap.Pop()
		res += int64(worker.cost)

		if l <= r {
			if worker.indx < l {
				heap.Push(Worker{l, costs[l]})
				l++
			} else {
				heap.Push(Worker{r, costs[r]})
				r--
			}
		}

	}
	return res
}
