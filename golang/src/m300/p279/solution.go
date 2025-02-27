package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func numSquares(n int) int {
	f := make([]int, n+1)
	s := make([]int, 0)

	for i := 0; i <= n; i++ {
		f[i] = i
	}

	for i := 1; ; i++ {
		x := i * i
		if x == n {
			return 1
		} else if x > n {
			break
		}
		s = append(s, x)
		f[x] = 1
	}

	for i := 2; i <= n; i++ {
		for _, k := range s {
			if k > i {
				break
			}
			f[i] = min(f[i], f[i-k]+1)
		}
	}
	return f[n]
}

func main() {
	r := numSquares(13)
	fmt.Println(r)
}
