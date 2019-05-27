package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	var arr = make([]int, 26)

	for i := 0; i < len(s); i++ {
		arr[s[i] - 'a']++
	}

	b := true
	for _, count := range arr {
		if count > 0 && count < k {
			b = false
		}
	}
	if b {
		return len(s)
	}

	pre, r := 0, 0
	for i := 0; i < len(s); i++ {
		if arr[s[i] - 'a'] < k {
			r = max(r, longestSubstring(s[pre:i], k))
			pre = i + 1
		}
	}
	r = max(r, longestSubstring(s[pre:], k))
	return r
}

func main() {
	r := longestSubstring("ababc", 2)
	fmt.Println(r)
}
