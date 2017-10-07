package main

import (
	"fmt"
	"p295"
	"p316"
	"p363"
	"p367"
	"p468"
	"p530"
)

func run540() {
	var x = []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(x))
}

func run316() {
	var s = "bcabc"
	fmt.Println(p316.RemoveDuplicateLetters(s))
}

func run434() {
	var s = "Hello, my name is John"
	fmt.Println(countSegments(s))
}

func run295() {
	var obj = p295.Constructor()
	_ = obj
}

func run530() {
	var n1 = p530.MakeTreeNode(1)
	var n2 = p530.MakeTreeNode(2)
	n2.Left = n1
	var n3 = p530.MakeTreeNode(4)
	var n4 = p530.MakeTreeNode(3)
	n4.Left = n2
	n4.Right = n3
	fmt.Println(p530.GetMinimumDifference(n4))
}

func run363() {
	var matrix = [][]int{
		[]int{5, -4, -3, 4},
		[]int{-3, -4, 4, 5},
		[]int{5, 1, 5, -4},
	}
	fmt.Println(p363.MaxSumSubmatrix(matrix, -2))

	var s = []int{1}
	_ = append(s, 4)
}

func run367() {
	fmt.Println(p367.IsPerfectSquare(168))
}

func run468() {
	fmt.Println(p468.ValidIPAddress("192.168.0.1"))
}

func main() {
	run468()
}
