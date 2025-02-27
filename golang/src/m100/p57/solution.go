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

func insert(intervals [][]int, newInterval []int) (r [][]int) {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	lo, hi := 0, len(intervals)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if newInterval[0] >= intervals[mid][0] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	i, j := lo-1, lo-1
	for ; i >= 0; i-- {
		if intervals[i][1] < newInterval[0] {
			break
		}
		newInterval[0] = intervals[i][0]
	}

	j = max(0, j)
	for ; j < len(intervals); j++ {
		if intervals[j][0] > newInterval[1] {
			break
		}
		newInterval[1] = max(newInterval[1], intervals[j][1])
	}

	r = append(r, intervals[:i+1]...)
	r = append(r, newInterval)
	r = append(r, intervals[j:]...)
	return
}

func main() {
	r := insert([][]int{
		{1, 5},
	}, []int{0, 0})
	fmt.Println(r)
}
