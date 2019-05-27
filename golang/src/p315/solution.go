package main

import (
	"fmt"
)

type BSTNode struct {
	val         int
	leftsum     int
	count       int
	left, right *BSTNode
}

type BST struct {
	root *BSTNode
}

func (t *BST) add(v int) int {
	if t.root == nil {
		t.root = &BSTNode{v, 0, 1, nil, nil}
		return 0
	} else if v == t.root.val {
		t.root.count++
		return 0
	}

	var parent, node *BSTNode
	parent, node, c := nil, t.root, 0
	for node != nil && node.val != v {
		if v < node.val {
			parent, node = node, node.left
			parent.leftsum++
		} else {
			c += node.count + node.leftsum
			parent, node = node, node.right
		}
	}
	if node != nil {
		c += node.leftsum
		node.count++
	} else if v < parent.val {
		parent.left = &BSTNode{v, 0, 1, nil, nil}
	} else {
		parent.right = &BSTNode{v, 0, 1, nil, nil}
	}
	return c
}

func countSmaller(nums []int) []int {
	r := make([]int, len(nums))

	t := &BST{}
	for i := len(nums) - 1; i >= 0; i-- {
		r[i] = t.add(nums[i])
	}

	return r
}

func main() {
	r := countSmaller([]int{5, 9, 6, 9, 1})
	fmt.Println(r)
}
