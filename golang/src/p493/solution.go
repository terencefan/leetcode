package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
	count int
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil, 1}
}

type BST struct {
	root *TreeNode
}

func (t *BST) Add(x int) {
	if t.root == nil {
		t.root = NewTreeNode(x)
		return
	}
	var node = t.root
	for {
		node.count++
		if x < node.val {
			if node.left == nil {
				node.left = NewTreeNode(x)
				break
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				node.right = NewTreeNode(x)
				break
			} else {
				node = node.right
			}
		}
	}
}

func (t *BST) countMoreThan(x int) (r int) {
	if t.root == nil {
		return 0
	}

	node := t.root
	for node != nil {
		if x >= node.val {
			node = node.right
		} else {
			r += 1
			if node.right != nil {
				r += node.right.count
			}
			node = node.left
		}
	}
	return
}

func NewTree() *BST {
	return &BST{}
}

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	currentMax := -1

	dp, t := 0, NewTree()
	for _, num := range nums {
		if num > currentMax {
			currentMax = num
		}
		if currentMax > num*2 {
			dp += t.countMoreThan(num * 2)
		}
		t.Add(num)
	}
	return dp
}

func main() {
	r := reversePairs([]int{1, 3, 2, 3, 1})
	fmt.Println(r)
}
