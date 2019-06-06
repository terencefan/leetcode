package main

func min(a, b int) int {
	if a > b || a < 0 {
		return b
	} else {
		return a
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func shortestDistance(words []string, word1 string, word2 string) int {
	p1, p2, d := -1, -1, -1

	for index, word := range words {
		if word == word1 {
			p1 = index
		} else if word == word2 {
			p2 = index
		}
		if p1 >= 0 && p2 >= 0 {
			d = min(d, abs(p1-p2))
		}
	}
	return d
}
