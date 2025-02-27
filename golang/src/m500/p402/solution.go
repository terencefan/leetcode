package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	var i, s = 0, make([]byte, 0)

	for i = 0; i < len(num) && k > 0; i++ {
		for len(s) > 0 && k > 0 {
			if num[i] < s[len(s)-1] {
				s = s[:len(s)-1]
				k--
			} else {
				break
			}
		}
		s = append(s, num[i])
	}
	s = append(s, num[i:]...)
	for k > 0 {
		s = s[:len(s)-1]
		k--
	}

	leadingZeros := 0
	for i := 0; i < len(s) && s[i] == '0'; i++ {
		leadingZeros++
	}
	if leadingZeros == len(s) {
		return "0"
	} else {
		return string(s[leadingZeros:])
	}
}

func main() {
	r := removeKdigits("10100", 1)
	fmt.Println(r)
}
