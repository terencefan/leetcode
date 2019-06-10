package main

import (
	"fmt"
)

func reverse(bytes []byte, l, r int) {
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
}

func reverseWords(s string) string {
	bytes := []byte(s)

	l, r := 0, 0
	for i := range bytes {
		c := bytes[i]
		if c == ' ' {
			reverse(bytes, l, r-1)
			l = i + 1
		}
		r++
	}
	reverse(bytes, l, r-1)
	return string(bytes)
}

func main() {
	fmt.Println(reverseWords("Let's find something new"))
}
