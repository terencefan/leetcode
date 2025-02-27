package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	m := 26
	for n > m {
		n -= m
		m *= 26
	}

	n -= 1

	bytes := make([]byte, 0)
	for m > 1 {
		m /= 26
		bytes = append(bytes, byte(n/m)+'A')
		n %= m
	}
	return string(bytes)
}

func main() {
	fmt.Println(convertToTitle(26))
}
