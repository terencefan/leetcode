package main

import (
	"fmt"
)

func test(s string, num, k int) bool {
	m := make(map[byte]int)
	for i := 0; i < num; i++ {
		m[s[i]]++
	}
	if len(m) <= k {
		return true
	}

	for i := num; i < len(s); i++ {
		l, r := s[i-num], s[i]
		m[l]--
		m[r]++
		if m[l] == 0 {
			delete(m, l)
		}
		if len(m) <= k {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) < k {
		return len(s)
	}

	l, r := k, len(s)
	for l <= r {
		m := l + (r-l)/2
		if test(s, m, k) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l - 1
}

func main() {
	r := lengthOfLongestSubstringKDistinct("ecexa", 2)
	fmt.Println(r)
}
