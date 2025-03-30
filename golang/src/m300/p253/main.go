package p253

import (
	"container/heap"
	"sort"
)

type Interval struct {
	start, end int
}

type IntervalHeap []Interval

// Len implements heap.Interface.
func (h *IntervalHeap) Len() int {
	return len(*h)
}

// Less implements heap.Interface.
func (h *IntervalHeap) Less(i int, j int) bool {
	return (*h)[i].end < (*h)[j].end
}

// Pop implements heap.Interface.
func (h *IntervalHeap) Pop() any {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

// Push implements heap.Interface.
func (h *IntervalHeap) Push(x any) {
	*h = append(*h, x.(Interval))
}

// Swap implements heap.Interface.
func (h *IntervalHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func minMeetingRooms(inputs [][]int) int {
	intervals := make([]Interval, len(inputs))

	for i, arr := range inputs {
		intervals[i] = Interval{arr[0], arr[1]}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	h := &IntervalHeap{}
	heap.Init(h)

	var rooms = 0
	for _, interval := range intervals {
		for h.Len() > 0 && (*h)[0].end <= interval.start {
			heap.Pop(h)
		}
		heap.Push(h, interval)
		rooms = max(rooms, h.Len())
	}
	return rooms
}
