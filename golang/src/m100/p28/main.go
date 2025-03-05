package main

func next(needle string) []int {
	next := make([]int, len(needle))
	next[0] = -1
	for i := 1; i < len(needle); i++ {
		k := next[i-1]
		for k >= 0 && needle[k+1] != needle[i] {
			k = next[k]
		}
		if needle[k+1] == needle[i] {
			next[i] = k + 1
		} else {
			next[i] = -1
		}
	}
	return next
}

func kmp(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := next(needle)

	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - len(needle)
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = next[j-1] + 1
			}
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	return kmp(haystack, needle)
}
