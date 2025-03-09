package p56

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	s, t := intervals[0][0], intervals[0][1]

	var r = make([][]int, 0)
	for i := range len(intervals) {
		if i == 0 {
			continue
		}

		si, ti := intervals[i][0], intervals[i][1]
		if si > t {
			r = append(r, []int{s, t})
			s, t = si, ti
		} else {
			if si < s {
				s = si
			}
			if ti > t {
				t = ti
			}
		}
	}
	r = append(r, []int{s, t})
	return r
}
