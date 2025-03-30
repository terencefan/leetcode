package p2931

import "container/heap"

type ShopVal struct {
	shop int
	val  int
}

type ShopValHeap []ShopVal

// Len implements heap.Interface.
func (s *ShopValHeap) Len() int {
	return len(*s)
}

// Less implements heap.Interface.
func (s *ShopValHeap) Less(i int, j int) bool {
	return (*s)[i].val < (*s)[j].val
}

// Pop implements heap.Interface.
func (s *ShopValHeap) Pop() any {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

// Push implements heap.Interface.
func (s *ShopValHeap) Push(x any) {
	*s = append(*s, x.(ShopVal))
}

// Swap implements heap.Interface.
func (s *ShopValHeap) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func maxSpending(values [][]int) int64 {
	h := &ShopValHeap{}

	for shop, shopVals := range values {
		if len(shopVals) > 0 {
			heap.Push(h, ShopVal{shop, shopVals[len(shopVals)-1]})
			values[shop] = shopVals[:len(shopVals)-1]
		}
	}

	r, d := int64(0), 1
	for h.Len() > 0 {
		shopVal := heap.Pop(h).(ShopVal)
		r += int64(d) * int64(shopVal.val)
		d++

		shop, shopValCount := shopVal.shop, len(values[shopVal.shop])
		if shopValCount > 0 {
			heap.Push(h, ShopVal{shop, values[shop][shopValCount-1]})
			values[shop] = values[shop][:shopValCount-1]
		}
	}
	return r
}
