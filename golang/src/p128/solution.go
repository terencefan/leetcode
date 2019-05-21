package main

import (
	"fmt"
)

type node struct {
	size int
}

func longestConsecutive(nums []int) (r int) {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)
	for _, num := range nums {
		m[num] = 1
	}

	for _, num := range nums {
		p, q, k := num, 0, 1
		for {
			if v, ok := m[p]; !ok {
				break
			} else {
				p += v
				k += v
			}
		}
		k--

		p, q = num, k
		for q > 1 {
			if v, ok := m[p]; ok {
				fmt.Println(p, q)
				m[p] = q
				p += v
				q -= v
			}
		}

		if k > r {
			r = k
		}
	}

	return r
}

func main() {
	r := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(r)
}
