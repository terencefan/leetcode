package p395

func longestSubstring(s string, k int) int {
	var m = make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	res := 0
	for maxCount := 1; maxCount <= len(m); maxCount++ {
		l, countMap, countK := 0, make(map[byte]int), 0
		for r, c := range s {
			countMap[byte(c)]++
			if countMap[byte(c)] == k {
				countK++
			}

			for len(countMap) > maxCount {
				if countMap[s[l]] == k {
					countK--
				}
				countMap[s[l]]--
				if countMap[s[l]] == 0 {
					delete(countMap, s[l])
				}
				l++
			}

			if countK == maxCount {
				res = max(res, r-l+1)
			}
		}
	}
	return res
}
