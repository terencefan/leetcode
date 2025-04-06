package p347

type NumCount struct {
	num, count int
}

type MinHeapK struct {
	data []NumCount
	size int
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func (h *MinHeapK) Push(num, count int) {
	if len(h.data) < h.size {
		h.data = append(h.data, NumCount{num: num, count: count})
		i := len(h.data) - 1
		for i > 0 {
			p := parent(i)
			if h.data[p].count <= h.data[i].count {
				break
			}
			h.data[p], h.data[i] = h.data[i], h.data[p]
			i = p
		}
	} else if count <= h.data[0].count {
		return
	} else {
		h.data[0] = NumCount{num: num, count: count}
		i := 0
		for {
			l, r := children(i)
			if l >= len(h.data) {
				break
			} else if r >= len(h.data) {
				if h.data[l].count < h.data[i].count {
					h.data[l], h.data[i] = h.data[i], h.data[l]
					i = l
				} else {
					break
				}
			} else {
				if h.data[l].count >= h.data[r].count && h.data[r].count < h.data[i].count {
					h.data[r], h.data[i] = h.data[i], h.data[r]
					i = r
				} else if h.data[r].count >= h.data[l].count && h.data[l].count < h.data[i].count {
					h.data[l], h.data[i] = h.data[i], h.data[l]
					i = l
				} else {
					break
				}
			}
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	h := &MinHeapK{size: k}
	for num, count := range freq {
		h.Push(num, count)
	}

	r := make([]int, k)
	for i, item := range h.data {
		r[i] = item.num
	}
	return r
}
