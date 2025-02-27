package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	word  string
	count int
}

type Heap []*Node

func greater(w1, w2 string) bool {
	if len(w1) < len(w2) {
		return !greater(w2, w1)
	}
	for i := range w2 {
		if w1[i] < w2[i] {
			return false
		} else if w1[i] > w2[i] {
			return true
		}
	}
	return true
}

func (h Heap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return greater(h[i].word, h[j].word)
	} else {
		return h[i].count < h[j].count
	}
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h Heap) Len() int {
	return len(h)
}

func topKFrequent(words []string, k int) []string {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	h := make(Heap, 0)
	for word, count := range counts {
		heap.Push(&h, &Node{word, count})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	var r = make([]string, h.Len())
	for h.Len() > 0 {
		node := heap.Pop(&h).(*Node)
		r[h.Len()] = node.word
	}
	return r
}

func main() {
	r := topKFrequent(
		[]string{"i", "love", "leetcode", "i", "love", "coding"},
		3,
	)
	fmt.Println(r)
}
