package main

import (
	"fmt"
)

type Node struct {
	pos, repeat int
}

func decodeString(s string) string {
	return string(decode(s))
}

func decode(s string) []byte {
	var repeat = 0
	var r []byte

	var first, count = -1, 0
	for i, c := range s {
		if c >= '0' && c <= '9' && count == 0 {
			repeat *= 10
			repeat += int(c - '0')
		} else if c == '[' {
			if count == 0 {
				first = i
			}
			count++
		} else if c == ']' {
			count--
			if count == 0 {
				part := s[first+1:i]
				t := decode(part)
				for i := 0; i < repeat; i++ {
					r = append(r, t...)
				}
				repeat = 0
			}
		} else if repeat == 0 {
			r = append(r, byte(c))
		}
	}
	return r
}

func main() {
	r := decodeString("3[a]2[bc]")
	fmt.Println(string(r))
}
