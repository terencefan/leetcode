package p435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var r, k = 0, -1 << 31
	for _, interval := range intervals {
		if interval[0] < k {
			r++
		} else {
			k = interval[1]
		}
	}
	return k
}
