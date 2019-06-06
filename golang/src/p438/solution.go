package main

import (
	"fmt"
)

func ich(b byte) int {
	return int(b) - int('a')
}

func isEmpty(arr []int) bool {
	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	if len(p)-len(s) > 0 {
		return make([]int, 0)
	}

	arr := make([]int, 26)
	pivot := 0

	for i := 0; i < len(p); i++ {
		pi := ich(p[i])
		si := ich(s[i])

		pivot += pi - si
		arr[pi]++
		arr[si]--
	}

	r := make([]int, 0)
	for i := 0; i < len(s)-len(p); i++ {
		if pivot == 0 && isEmpty(arr) {
			r = append(r, i)
		}
		s1, s2 := ich(s[i]), ich(s[i+len(p)])

		pivot += s1 - s2
		arr[s2]--
		arr[s1]++
	}
	if pivot == 0 && isEmpty(arr) {
		r = append(r, len(s)-len(p))
	}

	return r
}

func main() {
	r := findAnagrams("baa", "aa")
	fmt.Println(r)
}
