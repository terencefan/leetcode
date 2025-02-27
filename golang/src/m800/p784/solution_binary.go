package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {

	isLower := func(b byte) bool {
		return b >= 'a' && b <= 'z'
	}

	isUpper := func(b byte) bool {
		return b >= 'A' && b <= 'Z'
	}

	n := 1
	for i := range S {
		if isLower(S[i]) || isUpper(S[i]) {
			n <<= 1
		}
	}

	return r
}

func main() {
	fmt.Println(letterCasePermutation("a1C2"))
}
