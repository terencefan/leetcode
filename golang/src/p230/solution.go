package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	cur := root

	s := make([]*TreeNode, 0)

	for cur != nil {
		s = append(s, cur)
		cur = cur.Left
	}

	for k > 1 {
		cur = s[len(s)-1]
		s = s[:len(s)-1]
		if cur.Right != nil {
			cur = cur.Right
			for cur != nil {
				s = append(s, cur)
				cur = cur.Left
			}
		}
		k--
	}

	if len(s) > 0 {
		return s[len(s)-1].Val
	} else {
		return 0
	}
}

func main() {
	r := kthSmallest(&TreeNode{1, nil, &TreeNode{2, nil, nil}}, 2)
	fmt.Println(r)
}
