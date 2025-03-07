package p274

func hIndex(citations []int) int {
	var n = len(citations)

	counting := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			counting[n]++
		} else {
			counting[c]++
		}
	}

	current := 0
	for i := n; i > 0; i-- {
		current += counting[i]
		if current >= i {
			return i
		}
	}
	return 0
}
