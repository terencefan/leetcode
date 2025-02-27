package main

import (
	"fmt"
)

func isHappy(n int) bool {
	visited := make(map[int]bool)

	for n > 0 {
		r, m := 0, n
		for m > 0 {
			r += (m % 10) * (m % 10)
			m = m / 10
		}
		if visited[r] {
			return false
		} else if r == 1 {
			return true
		} else {
			visited[r] = true
			n = r
		}
	}
	return false
}

func main() {
	r := isHappy(19)
	fmt.Println(r)
}
