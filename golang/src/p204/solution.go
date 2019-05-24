package main

import (
	"fmt"
)

func countPrimes1(n int) int {
	s := make([]int, 0)

	for x := 2; x < n; x++ {
		b := true
		for _, y := range s {
			if x % y == 0 {
				b = false
				break
			}
		}
		if b {
			s = append(s, x)
		}
	}
	return len(s)
}

func countPrimes(n int) (r int) {
	s := make([]bool, n)

	for x := 2; x * x < n; x++ {
		if !s[x] {
			for i := x * x; i < n; i += x {
				s[i] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !s[i] {
			r++
		}
	}
	return
}

func main() {
	r := countPrimes(10)
	fmt.Println(r)
}
