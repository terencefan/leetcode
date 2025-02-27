package main

import (
	"fmt"
)

var r []string

func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isChar(b byte) bool {
	return isLower(b) || isUpper(b)
}

func dfs(bytes []byte, pos int) {
	for pos < len(bytes) && !isChar(bytes[pos]) {
		pos++
	}

	if pos == len(bytes) {
		r = append(r, string(bytes))
		return
	}

	dfs(bytes, pos+1)
	if isLower(bytes[pos]) {
		bytes[pos] -= byte('a') - byte('A')
		dfs(bytes, pos+1)
		bytes[pos] += byte('a') - byte('A')
	} else if isUpper(bytes[pos]) {
		bytes[pos] += byte('a') - byte('A')
		dfs(bytes, pos+1)
		bytes[pos] -= byte('a') - byte('A')
	}
}

func letterCasePermutation_recursive(S string) []string {
	r = make([]string, 0)
	dfs([]byte(S), 0)
	return r
}

func main() {
	fmt.Println(letterCasePermutation("a1C2"))
}
