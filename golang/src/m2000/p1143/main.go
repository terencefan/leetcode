package p1143

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 {
		return 0
	}
	if len(text2) > len(text1) {
		text1, text2 = text2, text1
	}

	var previous = make([]int, len(text1))
	var current = make([]int, len(text1))

	for j := range len(text2) {
		if text1[0] == text2[j] {
			current[0] = 1
		}
		for i := 1; i < len(text1); i++ {
			if text1[i] == text2[j] {
				current[i] = previous[i-1] + 1
			} else {
				current[i] = max(previous[i], current[i-1])
			}
		}
		previous, current = current, previous
	}
	return previous[len(text1)-1]
}
