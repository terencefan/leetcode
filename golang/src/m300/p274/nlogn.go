package p274

func hIndex(citations []int) int {

	l, r := 0, len(citations)+1

	for l < r {
		mid := l + (r-l)/2

		count := 0
		for _, c := range citations {
			if c >= mid {
				count++
				if count >= mid {
					break
				}
			}
		}

		if count >= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}
