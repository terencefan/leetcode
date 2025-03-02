package p1657

func dist(word string) map[rune]int {
	var m = make(map[rune]int)
	for _, c := range word {
		m[c]++
	}
	return m
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var s1 = dist(word1)
	var s2 = dist(word2)

	if len(s1) != len(s2) {
		return false
	}

	var vals1 = make(map[int]int)
	var vals2 = make(map[int]int)

	for k, v1 := range s1 {
		vals1[v1]++
		if v2, ok := s2[k]; ok {
			vals2[v2]++
		} else {
			return false
		}
	}

	for k, v1 := range vals1 {
		if v2, ok := vals2[k]; !ok {
			return false
		} else if v1 != v2 {
			return false
		}
	}
	return true
}
