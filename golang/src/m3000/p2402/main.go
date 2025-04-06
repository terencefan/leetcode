package p2402

import (
	"container/heap"
	"sort"
)

type Meeting struct {
	room       int
	start, end int
}

type MeetingHeap []Meeting

// Len implements heap.Interface.
func (m *MeetingHeap) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MeetingHeap) Less(i int, j int) bool {
	if (*m)[i].end == (*m)[j].end {
		return (*m)[i].room < (*m)[j].room
	}
	return (*m)[i].end < (*m)[j].end
}

// Pop implements heap.Interface.
func (m *MeetingHeap) Pop() any {
	item := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return item
}

// Push implements heap.Interface.
func (m *MeetingHeap) Push(x any) {
	*m = append(*m, x.(Meeting))
}

// Swap implements heap.Interface.
func (m *MeetingHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

type RoomHeap []int

// Len implements heap.Interface.
func (r *RoomHeap) Len() int {
	return len(*r)
}

// Less implements heap.Interface.
func (r *RoomHeap) Less(i int, j int) bool {
	return (*r)[i] < (*r)[j]
}

// Pop implements heap.Interface.
func (r *RoomHeap) Pop() any {
	item := (*r)[len(*r)-1]
	*r = (*r)[:len(*r)-1]
	return item
}

// Push implements heap.Interface.
func (r *RoomHeap) Push(x any) {
	*r = append(*r, x.(int))
}

// Swap implements heap.Interface.
func (r *RoomHeap) Swap(i int, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func mostBooked(n int, meetings [][]int) int {

	mh := make(MeetingHeap, 0)
	rh := make(RoomHeap, n)

	for i := range n {
		rh[i] = i
	}
	heap.Init(&rh)

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	bm, ri, rc := make(map[int]int), 0, 0
	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]

		for len(mh) > 0 && mh[0].end <= start {
			m := heap.Pop(&mh)
			heap.Push(&rh, m.(Meeting).room)
		}

		if len(rh) == 0 {
			m := heap.Pop(&mh)
			heap.Push(&rh, m.(Meeting).room)
			duration := end - start
			start, end = m.(Meeting).end, m.(Meeting).end+duration
		}

		room := heap.Pop(&rh).(int)
		bm[room]++

		if bm[room] > rc {
			ri, rc = room, bm[room]
		} else if bm[room] == rc {
			ri = min(ri, room)
		}

		heap.Push(&mh, Meeting{room, start, end})
	}
	return ri
}
