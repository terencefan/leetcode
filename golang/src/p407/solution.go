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

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func build(height, width int) [][]int {
	r := make([][]int, height)
	for i := range r {
		r[i] = make([]int, width)
	}
	return r
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	height, width := len(heightMap), len(heightMap[0])

	var (
		top   = build(height, width)
		down  = build(height, width)
		left  = build(height, width)
		right = build(height, width)
	)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k := width - j - 1
			if j == 0 {
				left[i][j] = heightMap[i][j]
				right[i][k] = heightMap[i][k]
			} else {
				left[i][j] = max(heightMap[i][j], left[i][j-1])
				right[i][k] = max(heightMap[i][k], right[i][k+1])
			}
		}
	}

	for j := 0; j < width; j++ {
		for i := 0; i < height; i++ {
			k := height - i - 1
			if i == 0 {
				top[i][j] = heightMap[i][j]
				down[k][j] = heightMap[k][j]
			} else {
				top[i][j] = max(heightMap[i][j], top[i-1][j])
				down[k][j] = max(heightMap[k][j], down[k+1][j])
			}
		}
	}

	r := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			height := min(min(top[i][j], down[i][j]), min(left[i][j], right[i][j]))
			r += height - heightMap[i][j]
		}
	}
	return r
}

func main() {
	r := trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	})
	fmt.Println(r)
}
