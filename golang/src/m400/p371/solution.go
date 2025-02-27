package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for {
		t := a & b
		a = a ^ b
		if t == 0 {
			return a
		}
		b = t << 1
	}
}

func main() {
	fmt.Println(getSum(13, 3))
}
