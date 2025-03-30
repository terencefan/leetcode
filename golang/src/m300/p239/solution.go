package main

type ValIndex struct {
	val, idx int
}

type MonoQueue []ValIndex

func (q *MonoQueue) push(vi ValIndex) {
	for len(*q) > 0 && (*q)[len(*q)-1].val < vi.val {
		*q = (*q)[:len(*q)-1]
	}
	*q = append(*q, vi)
}

func (q *MonoQueue) pop(idx int) {
	if len(*q) > 0 && (*q)[0].idx == idx {
		*q = (*q)[1:]
	}
}

func (q *MonoQueue) first() int {
	return (*q)[0].val
}

func maxSlidingWindow(nums []int, k int) []int {
	var r = make([]int, 0)
	var q = make(MonoQueue, 0)

	for i := range k {
		q.push(ValIndex{nums[i], i})
	}

	r = append(r, q.first())
	for i := k; i < len(nums); i++ {
		q.pop(i - k)
		q.push(ValIndex{nums[i], i})
		r = append(r, q.first())
	}
	return r
}
