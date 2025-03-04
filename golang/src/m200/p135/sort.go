package p135

import "sort"

type IndexRating struct {
	index  int
	rating int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	}

	var m = make([]IndexRating, len(ratings))
	var candy = make([]int, len(ratings))

	for i, rating := range ratings {
		m[i] = IndexRating{i, rating}
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].rating < m[j].rating
	})

	for _, item := range m {
		index := item.index

		lm, rm := 0, 0
		if index > 0 && item.rating > ratings[index-1] {
			lm = candy[index-1]
		}
		if index < len(ratings)-1 && item.rating > ratings[index+1] {
			rm = candy[index+1]
		}
		candy[index] = max(lm, rm) + 1
	}

	r := 0
	for _, v := range candy {
		r += v
	}
	return r
}
