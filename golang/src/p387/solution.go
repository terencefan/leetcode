package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	var m = make([]int, 26)
	var n = make([]int, 26)

	for i := len(s) - 1; i >= 0; i-- {
		b := s[i] - 'a'
		m[b]++
		n[b] = i
	}

	for b, count := range m {
		if count != 1 {
			n[b] = -1
		}
	}

	r := len(s)
	for _, index := range n {
		if index >= 0 && index < r {
			r = index
		}
	}
	if r < len(s) {
		return r
	} else {
		return -1
	}
}

func main() {
	r := firstUniqChar("teetcode")
	fmt.Println(r)
}