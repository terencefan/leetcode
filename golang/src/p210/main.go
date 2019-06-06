package main

import (
	"fmt"
)

func main() {
	var a = make([][]int, 4)
	a[0] = []int{1, 0}
	a[1] = []int{2, 0}
	a[2] = []int{3, 1}
	a[3] = []int{3, 2}
	fmt.Println(FindOrder(4, a))
}
