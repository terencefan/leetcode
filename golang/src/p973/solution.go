package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	dist  int
	point []int
}

type PointHeap []*Point

func (h PointHeap) Len() int      { return len(h) }
func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) { *h = append(*h, x.(*Point)) }

func (h *PointHeap) Pop() (r interface{}) {
	var l = len(*h)
	r = (*h)[l-1]
	*h = (*h)[:l-1]
	return
}

func (h PointHeap) Less(i, j int) bool { return h[i].dist > h[j].dist }

func dist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}

func kClosest(points [][]int, K int) (r [][]int) {
	h := make(PointHeap, 0)
	for _, point := range points {
		heap.Push(&h, &Point{dist(point), point})
		if h.Len() > K {
			heap.Pop(&h)
		}
	}

	for _, p := range h {
		r = append(r, p.point)
	}
	return
}

func main() {
	r := kClosest([][]int{
		{1, 3},
		{-2, 2},
	}, 1)
	fmt.Println(r)
}
