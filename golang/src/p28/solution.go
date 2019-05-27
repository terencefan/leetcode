package main

import (
	"fmt"
)

func kmp(s, t string) int {
	if len(t) == 0 {
		return 0
	}

	next := make([]int, len(t))
	next[0] = -1
	for i := 1; i < len(t); i++ {
		k := next[i-1]
		for k >= 0 && t[k+1] != t[i] {
			k = next[k]
		}
		if t[k+1] == t[i] {
			next[i] = k + 1
		} else {
			next[i] = -1
		}
	}

	i, j := 0, 0
	for i < len(s) {
		if s[i] == t[j] {
			i++
			j++
			if j == len(t) {
				return i - len(t)
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

func index(a, b string) int {
	if len(b) == 0 {
		return 0
	} else if len(b) == 1 {
		for k := 0; k < len(a); k++ {
			if a[k] == b[0] {
				return k
			}
		}
	} else if len(a) == len(b) {
		if a == b {
			return 0
		} else {
			return -1
		}
	} else if len(a) < len(b) {
		return -1
	}
	return kmp(a, b)
}

func strStr(haystack string, needle string) int {
	return index(haystack, needle)
}

func main() {
	r := strStr("mississippi", "issip")
	fmt.Println(r)
}
