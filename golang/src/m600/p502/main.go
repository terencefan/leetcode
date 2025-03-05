package p502

import (
	"container/heap"
	"sort"
)

type Project struct {
	profit  int
	capital int
}

type ProjectHeap []Project

// Len implements heap.Interface.
func (p *ProjectHeap) Len() int {
	return len(*p)
}

// Less implements heap.Interface.
func (p *ProjectHeap) Less(i int, j int) bool {
	return (*p)[i].profit > (*p)[j].profit
}

// Pop implements heap.Interface.
func (p *ProjectHeap) Pop() any {
	if len(*p) == 0 {
		return nil
	}
	r := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return r
}

// Push implements heap.Interface.
func (p *ProjectHeap) Push(x any) {
	*p = append(*p, x.(Project))
}

// Swap implements heap.Interface.
func (p *ProjectHeap) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make(ProjectHeap, len(profits))
	for i := range profits {
		projects[i] = Project{profits[i], capital[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	h := &ProjectHeap{}
	heap.Init(h)

	j := 0
	for i := 0; i < k; i++ {
		for ; j < len(projects) && projects[j].capital <= w; j++ {
			heap.Push(h, projects[j])
		}
		if len(*h) == 0 {
			return w
		}
		project := heap.Pop(h)
		w += project.(Project).profit
	}
	return w
}
