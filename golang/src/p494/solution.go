package main

import (
	"fmt"
)

func findTargetSumWays1(nums []int, S int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	if S > s || (s+S)&1 == 1 {
		return 0
	}

	var m1, m2 = make(map[int]int), make(map[int]int)
	var ms = []map[int]int{m1, m2}

	m1[0] = 1
	for i, n := range nums {
		m1 = ms[i%2]
		m2 = ms[(i+1)%2]

		for k := range m2 {
			delete(m2, k)
		}

		for k, v := range m1 {
			m2[k+n] += v
			m2[k-n] += v
		}
	}

	m := ms[len(nums)%2]
	return m[S]
}

func findTargetSumWays(nums []int, S int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	if S > s || (s+S)&1 == 1 {
		return 0
	}

	var f = make([]int, s+1)
	f[0] = 1

	for _, n := range nums {
		for i := s; i >= n; i-- {
			f[i] += f[i-n]
		}
	}
	return f[(S+s)/2] // TODO but why?
}

func main() {
	r := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println(r)
}
