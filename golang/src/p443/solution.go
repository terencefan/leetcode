package main

import (
	"fmt"
)

func process(cur byte, chars []byte, count, pos int) int {
	if count > 1 {
		chars[pos] = cur
		pos++

		p := 1
		for p <= count {
			p *= 10
		}

		for p > 1 {
			p /= 10
			chars[pos] = byte(count/p + '0')
			count %= p
			pos++
		}
	} else if count > 0 {
		chars[pos] = cur
		pos++
	}
	return pos
}

func compress(chars []byte) int {
	cur, count, pos := byte('^'), 0, 0

	for _, char := range chars {
		if char != cur {
			pos = process(cur, chars, count, pos)
			count = 1
		} else {
			count++
		}
		cur = char
	}
	return process(cur, chars, count, pos)
}

func main() {
	fmt.Println(compress([]byte("abcccccccccc$$")))
}
