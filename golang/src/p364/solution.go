package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxDepth(nestedList []*NestedInteger) int {
	depth := 1
	for _, item := range nestedList {
		if !item.IsInteger() {
			depth = max(maxDepth(item.GetList())+1, depth)
		}
	}
	return depth
}

func count(nestedList []*NestedInteger, depth int) (r int) {
	for _, item := range nestedList {
		if item.IsInteger() {
			r += depth * item.GetInteger()
		} else {
			r += count(item.GetList(), depth-1)
		}
	}
	return
}

func depthSumInverse(nestedList []*NestedInteger) int {
	depth := maxDepth(nestedList)
	return count(nestedList, depth)
}

func main() {
	fmt.Println("vim-go")
}
