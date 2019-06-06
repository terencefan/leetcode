package main

import (
	"fmt"
)

func charAt(s1, s2 string, x int) byte {
	if x >= len(s1) {
		return s2[x-len(s1)]
	} else {
		return s1[x]
	}
}

func test(s1, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	i, j := 0, l1+l2-1
	for i < j {
		if charAt(s1, s2, i) != charAt(s1, s2, j) {
			return false
		}
		i++
		j--
	}
	return true
}

func palindromePairs1(words []string) [][]int {
	r := make([][]int, 0)

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && test(words[i], words[j]) {
				r = append(r, []int{i, j})
			}
		}
	}
	return r
}

func main() {
	r := palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"})
	fmt.Println(r)
}
