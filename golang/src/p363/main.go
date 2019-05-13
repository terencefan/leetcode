package main


import (
	"fmt"
)

func main() {
	var matrix = [][]int{
		{5, -4, -3, 4},
		{-3, -4, 4, 5},
		{5, 1, 5, -4},
	}
	r := MaxSumSubmatrix(matrix, -2)
	fmt.Println(r)
}

