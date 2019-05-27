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

func countLessEqualThan(arr []int, val int) int {
	if val >= arr[len(arr) - 1] {
		return len(arr)
	} else if val < arr[0] {
		return 0
	}

	lo, hi := 0, len(arr) - 1

	for lo < hi {
		mid := lo + (hi - lo) / 2
		if arr[mid] > val {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])

	lo, hi := matrix[0][0], matrix[m - 1][n - 1]
	for lo < hi {

		mid := lo + (hi - lo) / 2

		count := 0
		for i := 0; i < m; i++ {
			count += countLessEqualThan(matrix[i], mid)
		}

		if count >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	// fmt.Println(countLessEqualThan([]int{1,3,3,7,8}, 1))

	r := kthSmallest([][]int{
		{1,  5,  9},
		{10, 11, 13},
		{12, 13, 15},
	}, 2)
	fmt.Println(r)
}
