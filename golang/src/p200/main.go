package main

import (
	"fmt"
)

func main() {
	r := numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})
	fmt.Println(r)
}
