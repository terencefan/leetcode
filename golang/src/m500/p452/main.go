package p452

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	var boom = make([]bool, len(points))

	var k = 0
	for i := range points {
		if boom[i] {
			continue
		}
		k++

		end := points[i][1]

		for j := i; j < len(points); j++ {
			if points[j][0] > end {
				break
			}
			boom[j] = true
			end = min(end, points[j][1])
		}
	}
	return k
}
