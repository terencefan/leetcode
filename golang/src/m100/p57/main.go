package p57

func insert(intervals [][]int, newInterval []int) [][]int {
	s, t := newInterval[0], newInterval[1]
	inserted := false

	r := make([][]int, 0)
	for i := range len(intervals) {
		si, ti := intervals[i][0], intervals[i][1]
		if ti < s {
			r = append(r, []int{si, ti})
		} else if si > t {
			if !inserted {
				r = append(r, []int{s, t})
				inserted = true
			}
			r = append(r, []int{si, ti})
		} else {
			if ti > t {
				t = ti
			}
			if si < s {
				s = si
			}
		}
	}

	if !inserted {
		r = append(r, []int{s, t})
		inserted = true
	}
	return r
}
