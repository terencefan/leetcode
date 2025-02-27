package main

import (
	"fmt"
)

func main() {
	r := ShoppingOffers(
		[]int{2, 5},
		[][]int{
			{3, 0, 5},
			{1, 2, 10},
		},
		[]int{3, 2},
	)
	fmt.Println(r)
}
