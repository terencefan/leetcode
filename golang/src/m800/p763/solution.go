package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func partitionLabels(s string) (r []int) {
	if len(s) == 0 {
		return []int{}
	}

	last := make([]int, 26)

	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	next := 0
	for i := 0; i < len(s); i++ {
		next = max(next, last[s[i]-'a'])
		if i == next {
			r = append(r, i)
		}
	}

	for i := len(r) - 1; i > 0; i-- {
		r[i] -= r[i-1]
	}
	r[0] += 1
	return
}

func main() {
	r := partitionLabels("a")
	fmt.Println(r)
}
