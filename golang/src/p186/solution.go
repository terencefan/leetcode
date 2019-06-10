package main

import (
	"fmt"
)

func reverse(str []byte, l, r int) {
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
}

func reverseWords(str []byte) {
	reverse(str, 0, len(str)-1)

	l, r := 0, 0
	for i := range str {
		c := str[i]
		switch c {
		case ' ':
			reverse(str, l, r-1)
			l = i + 1
		}
		r++
	}
	reverse(str, l, r-1)
}

func main() {
	s := []byte("This is a problem")
	reverseWords(s)
	fmt.Println(string(s))
}
