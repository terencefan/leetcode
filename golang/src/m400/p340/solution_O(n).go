package main

import (
	"fmt"
)

func lengthOfLongestSubstringKDistinct(s string, k int) (r int) {
	i, j := 0, 0

	var m = make(map[byte]int)

	for j < len(s) {
		m[s[j]]++
		for len(m) > k {
			m[s[i]]--
			if m[s[i]] == 0 {
				delete(m, s[i])
			}
			i++
		}
		if j-i+1 > r {
			r = j - i + 1
		}
		j++
	}
	return
}

func main() {
	r := lengthOfLongestSubstringKDistinct("eceea", 2)
	fmt.Println(r)
}
