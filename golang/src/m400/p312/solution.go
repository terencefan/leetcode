package main

import (
	"fmt"
)

type balloons struct {
	nums []int
}

func (b *balloons) get(n int) int {
	if n == -1 {
		return 1
	} else if n > len(b.nums)-1 {
		return 1
	} else {
		return b.nums[n]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	b := &balloons{nums}

	f := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		f[i] = make([]int, len(nums))
		f[i][i] = b.get(i-1) * b.get(i) * b.get(i+1)
		if i+1 < len(nums) {
			l, m, n, r := b.get(i-1), b.get(i), b.get(i+1), b.get(i+2)
			f[i][i+1] = max(l*m*n+l*n*r, m*n*r+l*m*r)
		}
	}

	for i := 2; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			l, m, n, r := b.get(j-1), b.get(j), b.get(i+j), b.get(i+j+1)
			x := max(l*m*n+l*n*r, m*n*r+l*m*r)
			if i > 1 {
				x += f[j+1][j+i-1]
			}

			for k := j; k <= j+i; k++ {
				y := b.get(k) * l * r
				if k > j {
					y += f[j][k-1]
				}
				if k < j+i {
					y += f[k+1][j+i]
				}
				x = max(x, y)
			}
			f[j][j+i] = x
		}
	}

	return f[0][len(nums)-1]
}

func main() {
	r := maxCoins([]int{1, 4, 3, 2, 5})

	fmt.Println(r)
}
