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

type Node struct {
	pos, height int
}

func maxRectangle(dp []int) (r int) {
	var s = []Node{{-1,  0}}
	for i, height := range dp {
		for len(s) > 0 && s[len(s) - 1].height > height {
			node := s[len(s) - 1]
			s = s[:len(s) - 1]
			r = max(r, node.height * (i - s[len(s) - 1].pos - 1))
		}
		s = append(s, Node{i, height})
	}
	return
}

func maximalRectangle(matrix [][]byte) (r int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	dp := make([]int, len(matrix[0]) + 1)

	for i := range matrix {
		for k := range matrix[i] {
			if matrix[i][k] == '1' {
				dp[k]++
			} else {
				dp[k] = 0
			}
		}
		r = max(r, maxRectangle(dp))
	}
	return
}

func main() {
	r := maximalRectangle([][]byte{
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'},
	})
	fmt.Println(r)
}
