package p2561

import "sort"

func frequency(basket []int) map[int]int {
	var m = make(map[int]int)
	for _, fruit := range basket {
		m[fruit]++
	}
	return m
}

func minCost(basket1 []int, basket2 []int) int64 {
	f1, f2 := frequency(basket1), frequency(basket2)

	f := make(map[int]int)
	for k, v := range f1 {
		f[k] += v
	}
	for k, v := range f2 {
		f[k] += v
	}

	added, removed := make([]int, 0), make([]int, 0)

	minv := 1 << 31
	for k, v := range f {
		minv = min(minv, k)
		if v%2 == 1 {
			return -1
		}
		if f1[k] > v/2 {
			for range f1[k] - v/2 {
				removed = append(removed, k)
			}
		} else if f1[k] < v/2 {
			for range v/2 - f1[k] {
				added = append(added, k)
			}
		}
	}

	sort.Slice(removed, func(i, j int) bool {
		return removed[i] > removed[j]
	})
	sort.Slice(added, func(i, j int) bool {
		return added[i] < added[j]
	})

	r := int64(0)
	for i := range len(removed) {
		m := min(removed[i], added[i], minv*2)
		r += int64(m)
	}
	return r
}
