package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeAlternately(word1 string, word2 string) string {
	var r = make([]byte, 0)

	for i := 0; i < max(len(word1), len(word2)); i++ {
		if i < len(word1) {
			r = append(r, word1[i])
		}
		if i < len(word2) {
			r = append(r, word2[i])
		}
	}
	return string(r)
}

func main() {
	var r = mergeAlternately("abc", "pqr")
	fmt.Println(r)
}
