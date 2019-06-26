package main

import (
	"fmt"
)

func canCross(stones []int) bool {

	N := len(stones)
	if N == 0 {
		return true
	} else if N == 1 {
		return stones[0] == 0
	} else if !(stones[0] == 0 && stones[1] == 1) {
		return false
	}

	var f = make(map[int]map[int]bool, N)

	for i := 0; i < N; i++ {
		f[stones[i]] = make(map[int]bool)
	}
	f[1][1] = true

	var delta = []int{-1, 0, 1}

	for i, k := range stones {
		if i == 0 {
			continue
		}

		for step := range f[k] {
			for _, d := range delta {
				if step + d <= 0 {
					continue
				}
				next := k + step + d

				if _, ok := f[next]; ok {
					f[next][step + d] = true
				}
			}
		}
	}

	for _, k := range stones {
		fmt.Println(k, len(f[k]))
	}

	return len(f[stones[N - 1]]) > 0
}

func main() {
	var data = []int{0, 1}
	for i := 1; i < 1000; i++ {
		data = append(data, 2 * i)
	}
	r := canCross(data)
	fmt.Println(r)
}
