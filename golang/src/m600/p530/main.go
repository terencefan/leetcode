package main

import (
	"fmt"
)

func main() {
	var n1 = MakeTreeNode(1)
	var n2 = MakeTreeNode(2)
	n2.Left = n1
	var n3 = MakeTreeNode(4)
	var n4 = MakeTreeNode(3)
	n4.Left = n2
	n4.Right = n3

	r := GetMinimumDifference(n4)
	fmt.Println(r)
}
