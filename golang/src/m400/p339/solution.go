package main

import "fmt"

func count(nestedList []*NestedInteger, depth int) (r int) {
	for _, item := range nestedList {
		if item.IsInteger() {
			r += depth * item.GetInteger()
		} else {
			r += count(item.GetList(), depth+1)
		}
	}
	return
}

func depthSumInverse(nestedList []*NestedInteger) int {
	return count(nestedList, 1)
}

func main() {
	fmt.Println("vim-go")
}
