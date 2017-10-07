package p382

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) (s Solution) {
	s = Solution{Head: head}
	return
}

func (this *Solution) GetRandom() (res int) {
	var node = this.Head
	var i = 1

	for node != nil {
		if rand.Intn(i) == i-1 {
			res = node.Val
		}
		node = node.Next
		i++
	}
	return
}
