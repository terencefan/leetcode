package main

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	x := 1
	for ;x < n; x *= 3 {}
	return x == n
}

func main() {
	r := isPowerOfThree(27)
	fmt.Println(r)
}
