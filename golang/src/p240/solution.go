package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, 0

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	for i < m-1 && matrix[i][j] < target {
		i++
	}
	for j < n-1 && matrix[i][j] < target {
		j++
	}

	s := true
	for matrix[i][j] != target && s {
		fmt.Println(i, j)
		s = false
		for i > 0 && matrix[i][j] > target {
			i--
			s = true
		}
		for j < n-1 && matrix[i][j] < target {
			j++
			s = true
		}
	}
	return s
}

func main() {
	r := searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 4)
	fmt.Println(r)
}
